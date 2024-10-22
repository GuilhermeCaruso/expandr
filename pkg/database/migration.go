package database

import (
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

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
	Hash   string         `json:"hash"`
	Status MigratorStatus `json:"status"`
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

func NewMigrator() *Migrator {
	migrator := new(Migrator)
	return migrator
}

func (m *Migrator) RegisterMigration(migration Migration) {
	m.migrations = append(m.migrations, migration)
}

func (m *Migrator) Execute(db *Database, cmd MigratorCommand) {
	m.db = db

	switch cmd {
	case INIT:
		m.InitDatabaseMigrator()
	case UP:
		m.RunUp()
	case DOWN:
		break

	}
}

func (m *Migrator) RunUp() {
	lastMigration := new(MigratorTable)

	if err := m.db.Conn.Last(lastMigration).Error; err != nil {
		if gorm.ErrRecordNotFound != err {
			log.Fatalf("something went wrong searching last migration. %q", err)
		}
	}

	migrationsToUp := make([]Migration, 0)

	if lastMigration.ID == 0 {
		migrationsToUp = m.migrations
	}

	m.db.Conn.Transaction(func(tx *gorm.DB) error {
		for x := 0; x < len(migrationsToUp); x++ {
			mtu := migrationsToUp[x]
			if err := mtu.Up(context.Background(), tx); err != nil {
				log.Fatalf("fail to run migration %v. err:%q", mtu.Name, err)
				tx.Transaction(func(tx2 *gorm.DB) error {
					mt := MigratorTable{Hash: mtu.Name, Status: FAILED}
					return tx2.Create(&mt).Error
				})
				return err
			} else {
				log.Printf("migration %v was executed\n", mtu.Name)
				tx.Transaction(func(tx3 *gorm.DB) error {
					mt := MigratorTable{Hash: mtu.Name, Status: SUCCESS}
					return tx3.Create(&mt).Error
				})

			}
		}
		return nil
	})

	fmt.Println(migrationsToUp)
}

func (m *Migrator) InitDatabaseMigrator() {
	if m.db.Conn.Migrator().HasTable(&MigratorTable{}) {
		log.Println("database already initialized")
		return
	}

	if err := m.db.Conn.Migrator().CreateTable(&MigratorTable{}); err != nil {
		log.Fatalf("something went wrong trying to setup database. %q", err)
		return
	}

	log.Println("database initialized successfully")

}
