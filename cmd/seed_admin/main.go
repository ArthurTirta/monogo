package main

import (
    "context"
    "errors"
    "flag"
    "fmt"
    "log"

    "github.com/ArthurTirta/monogo/config"
    adminrepoimpl "github.com/ArthurTirta/monogo/internal/repository/admin/implementation"
    "github.com/ArthurTirta/monogo/internal/entity"
    databasehelper "github.com/ArthurTirta/monogo/pkg/helper/database"
    passwordhelper "github.com/ArthurTirta/monogo/pkg/helper/password"
    "gorm.io/gorm"
)

func main() {
    email := flag.String("email", "admin@admin.com", "Admin email")
    password := flag.String("password", "admin123", "Admin password")
    name := flag.String("name", "Administrator", "Admin name")
    flag.Parse()

    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("load config: %v", err)
    }

    db, err := databasehelper.NewGormDB(context.Background(), &cfg.DatabaseConfig)
    if err != nil {
        log.Fatalf("db connect: %v", err)
    }

    repo := adminrepoimpl.NewAdminRepository(db)

    // check existing
    _, err = repo.GetByEmail(context.Background(), *email)
    if err == nil {
        fmt.Printf("admin already exists: %s\n", *email)
        return
    }
    if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        log.Fatalf("failed checking admin: %v", err)
    }

    hash, err := passwordhelper.HashPassword(*password)
    if err != nil {
        log.Fatalf("hash password: %v", err)
    }

    adm := &entity.Admin{
        Name:     *name,
        Email:    *email,
        Password: hash,
    }

    out, err := repo.Create(context.Background(), adm)
    if err != nil {
        log.Fatalf("create admin: %v", err)
    }

    fmt.Printf("admin created: %s (id=%s)\n", out.Email, out.ID.String())
}
package main

import (
    "context"
    "errors"
    "flag"
    "fmt"
    "log"

    "github.com/ArthurTirta/monogo/config"
    adminrepoimpl "github.com/ArthurTirta/monogo/internal/repository/admin/implementation"
    "github.com/ArthurTirta/monogo/internal/entity"
    databasehelper "github.com/ArthurTirta/monogo/pkg/helper/database"
    passwordhelper "github.com/ArthurTirta/monogo/pkg/helper/password"
    "gorm.io/gorm"
)

func main() {
    email := flag.String("email", "admin@admin.com", "admin email")
    password := flag.String("password", "admin123", "admin password")
    name := flag.String("name", "Admin", "admin name")
    flag.Parse()

    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("load config: %v", err)
    }

    ctx := context.Background()
    db, err := databasehelper.NewGormDB(ctx, &cfg.DatabaseConfig)
    if err != nil {
        log.Fatalf("connect db: %v", err)
    }

    repo := adminrepoimpl.NewAdminRepository(db)

    // check existing
    _, err = repo.GetByEmail(ctx, *email)
    if err == nil {
        fmt.Printf("admin with email %s already exists, skipping\n", *email)
        return
    }
    if !errors.Is(err, gorm.ErrRecordNotFound) {
        log.Fatalf("error checking admin: %v", err)
    }

    hash, err := passwordhelper.HashPassword(*password)
    if err != nil {
        log.Fatalf("hash password: %v", err)
    }

    admin := entity.Admin{
        Name:     *name,
        Email:    *email,
        Password: hash,
    }

    out, err := repo.Create(ctx, &admin)
    if err != nil {
        log.Fatalf("create admin: %v", err)
    }

    fmt.Printf("admin created: id=%s email=%s\n", out.ID.String(), out.Email)
}
