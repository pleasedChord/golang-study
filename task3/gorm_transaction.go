package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Accounts struct {
	Id      string
	Balance float64
}

/*
accounts 表（包含字段 id 主键， balance 账户余额）
*/
type Transactions struct {
	Id              int
	From_account_id string
	To_account_id   string
	Amount          float64
}

func initAccount(db *gorm.DB) {
	//清除历史数据
	if db.Migrator().HasTable(&Accounts{}) {
		if err := db.Migrator().DropTable(&Accounts{}); err != nil {
			fmt.Println("Accounts表删除失败")
			return
		}
	}

	if db.Migrator().HasTable(&Transactions{}) {
		if err := db.Migrator().DropTable(&Transactions{}); err != nil {
			fmt.Println("Transactions表删除失败")
			return
		}
	}

	//创建表
	db.AutoMigrate(&Accounts{}, &Transactions{})

	//插入数据
	db.Create(&Accounts{
		Id:      "A",
		Balance: 999,
	})

	db.Create(&Accounts{
		Id:      "B",
		Balance: 0,
	})
}

func Transfer(db *gorm.DB, from, to string, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("转账金额必须大于0")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		//1.校验转出账户
		var fromAccount Accounts
		if err := tx.First(&fromAccount, "id = ?", from).Error; err != nil {
			return fmt.Errorf("转出账户不存在")
		}
		if fromAccount.Balance < amount {
			return fmt.Errorf("余额不足")
		}

		//2.校验转入账户
		var toAccount Accounts
		if err := tx.First(&toAccount, "id = ?", to).Error; err != nil {
			return fmt.Errorf("转入账户不存在")
		}

		//3.转出
		result := tx.Model(&Accounts{}).
			Where("id = ? and balance >= ?", from, amount).
			UpdateColumn("balance", gorm.Expr("balance - ?", amount))
		if result.Error != nil {
			return fmt.Errorf("转出失败:%v", result.Error)
		}
		if result.RowsAffected == 0 {
			return fmt.Errorf("余额不足或转出账户不存在")
		}

		//4.转入
		result = tx.Model(&Accounts{}).
			Where("id = ?", to).
			UpdateColumn("balance", gorm.Expr("balance + ?", amount))
		if result.Error != nil {
			return fmt.Errorf("转入失败:%v", result.Error)
		}

		//5.记录表记录并返回
		return tx.Create(&Transactions{
			From_account_id: from,
			To_account_id:   to,
			Amount:          amount,
		}).Error
	})
}

func PrintAccount(db *gorm.DB) {
	//1.查询客户表
	var accounts []Accounts
	db.Find(&accounts)
	for _, account := range accounts {
		fmt.Println("账户", account.Id, "的余额：", account.Balance)
	}

	//2.查询转账表
	var transactions []Transactions
	db.Find(&transactions)
	for _, trans := range transactions {
		fmt.Println("交易记录：", trans)
	}
}

func main() {
	//1.连接数据库
	db, err := gorm.Open(sqlite.Open("account.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接db失败")
	}

	//2.初始化表
	initAccount(db)
	PrintAccount(db)

	//3.转账
	err = Transfer(db, "A", "B", 100)
	if err != nil {
		fmt.Printf("转账失败: %v", err)
		fmt.Println()
	} else {
		fmt.Println("转账成功")
	}

	//4.打印结果
	PrintAccount(db)
}
