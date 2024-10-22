package database

import (
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

type MigratorType string

type MigratorStatus string

const (
	FAILED  MigratorStatus = "failed"
	SUCCESS MigratorStatus = "success"
)

type MigratorCommand string

const (
	INIT MigratorCommand = "init"
	UP   MigratorCommand = "up"
	DOWN MigratorCommand = "down"
)

type MigratorTable struct {
	gorm.Model
	Hash   string          `json:"hash"`
	Status MigratorStatus  `json:"status"`
	Type   MigratorCommand `json:"type"`
}

func (MigratorTable) TableName() string {
	return "migrator"
}

type MigrationFunction func(ctx context.Context, db *gorm.DB) error

type Migration struct {
	Name string
	Up   MigrationFunction
	Down MigrationFunction
}

type Migrator struct {
	db         *Database
	migrations []Migration
}

type MigratorParams struct {
	Cmd         MigratorCommand
	VersionHash *string
}

func NewMigrator() *Migrator {
	migrator := new(Migrator)
	return migrator
}

func (m *Migrator) RegisterMigration(migration Migration) {
	m.migrations = append(m.migrations, migration)
}

func (m *Migrator) Execute(db *Database, params MigratorParams) {
	m.db = db

	switch params.Cmd {
	case INIT:
		m.InitDatabaseMigrator()
	case UP:
		m.RunUp()
	case DOWN:
		if params.VersionHash != nil {
			m.RunDown(*params.VersionHash)
		} else {
			log.Fatalf("version to down is required.")

		}

	}
}

func (m *Migrator) checkMigratorTable() error {
	if !m.db.Conn.Migrator().HasTable(&MigratorTable{}) {
		log.Fatal("initialize migration controller first.")
		return errors.New("failed to verify database")
	}
	return nil
}

func (m *Migrator) getLastMigration() (*MigratorTable, error) {
	lastMigration := new(MigratorTable)

	if err := m.db.Conn.Where("status = ?", SUCCESS).Last(lastMigration).Error; err != nil {
		if gorm.ErrRecordNotFound != err {
			return nil, err
		}
	}

	return lastMigration, nil
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func (m *Migrator) Exec(migrations []Migration, action MigratorCommand) error {
	if len(migrations) == 0 {
		log.Println("no migration to perform")
		return nil
	}
	return m.db.Conn.Transaction(func(tx *gorm.DB) error {
		for x := 0; x < len(migrations); x++ {
			mtu := migrations[x]
			if err := mtu.Down(context.Background(), tx); err != nil {
				log.Fatalf("fail to run migration %v. err:%q", mtu.Name, err)
				tx.Transaction(func(tx2 *gorm.DB) error {
					mt := MigratorTable{Hash: mtu.Name, Status: FAILED, Type: action}
					return tx2.Create(&mt).Error
				})
				return err
			} else {
				log.Printf("migration %v was executed\n", mtu.Name)
				tx.Transaction(func(tx3 *gorm.DB) error {
					mt := MigratorTable{Hash: mtu.Name, Status: SUCCESS, Type: action}
					return tx3.Create(&mt).Error
				})

			}
		}
		return nil
	})
}

func (m *Migrator) RunDown(to string) {
	if err := m.checkMigratorTable(); err != nil {
		return
	}

	generalMigrations := m.migrations
	reverse(generalMigrations)

	lastMigrationsID := 0
	for x := 0; x < len(generalMigrations); x++ {
		if generalMigrations[x].Name == to {
			lastMigrationsID = x
			break
		}
	}

	migrationsToDown := generalMigrations[:lastMigrationsID]

	if err := m.Exec(migrationsToDown, DOWN); err != nil {
		log.Fatalf("some migrations had problems.err=%q", err)
		return
	}

	log.Println("all migrations were performed.")
}

func (m *Migrator) RunUp() {

	if err := m.checkMigratorTable(); err != nil {
		return
	}

	lastMigration, err := m.getLastMigration()

	if err != nil {
		log.Fatalf("something went wrong searching last migration. %q", err)
	}

	migrationsToUp := make([]Migration, 0)

	if lastMigration.ID == 0 {
		migrationsToUp = m.migrations
	} else {
		initialCursor := 0
		for x := 0; x < len(m.migrations); x++ {
			if m.migrations[x].Name == lastMigration.Hash {
				initialCursor = x
				break
			}
		}
		migrationsToUp = m.migrations[initialCursor+1:]
	}

	if err := m.Exec(migrationsToUp, UP); err != nil {
		log.Fatalf("some migrations had problems.err=%q", err)
		return
	}

	log.Println("all migrations were performed.")
}

func (m *Migrator) InitDatabaseMigrator() {
	if m.db.Conn.Migrator().HasTable(&MigratorTable{}) {
		log.Println("database already initialized.")
		return
	}

	if err := m.db.Conn.Migrator().CreateTable(&MigratorTable{}); err != nil {
		log.Fatalf("something went wrong trying to setup database. %q", err)
		return
	}

	log.Println("database initialized successfully.")

}
