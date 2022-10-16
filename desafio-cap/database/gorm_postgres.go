package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormPostgres struct {
	db *gorm.DB
}

func (gm *gormPostgres) connectService(config *OptionsDBClient) error {
	i := 0
	fmt.Println("DB URL - ", config.URL)
	for {
		db, err := gorm.Open(postgres.Open(config.URL), &gorm.Config{})
		if err != nil {
			if i >= 4 {
				return err
			} else {
				i++
			}
			time.Sleep(time.Second * 2)

		} else {
			gm.db = db
			log.Println("Connect on DataBase")
			break
		}

	}

	return nil
}

func (gm *gormPostgres) CreateSequence(sequence *Sequence) error {

	return gm.db.Clauses(clause.OnConflict{DoNothing: true}).Create(sequence).Error
}

func (gm *gormPostgres) GetInfoSequences() (*InfoSequence, error) {
	type result struct {
		Total int
	}
	infoSequence := &InfoSequence{}
	totalValid := &result{}
	totalInvalid := &result{}
	err := gm.db.Model(&Sequence{}).Select("is_valid, count(letters) as total").Group("is_valid").Having("is_valid = ?", true).Find(&totalValid).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}

	err = gm.db.Model(&Sequence{}).Select("is_valid, count(letters) as total").Group("is_valid").Having("is_valid = ?", false).Find(&totalInvalid).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	infoSequence.Valids = totalValid.Total
	infoSequence.Invalids = totalInvalid.Total
	infoSequence.Ratio = float32(infoSequence.Valids) / float32((infoSequence.Valids + infoSequence.Invalids))
	return infoSequence, nil

}
