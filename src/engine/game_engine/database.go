package game_engine

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strconv"
)

func newDatabase() (*sql.DB, error) {
	dbPath := os.Getenv("ROOTIQUE") + "/common/database/db.sqlite3"
	return sql.Open("sqlite3", dbPath)
}

func loadTerrains(db *sql.DB) ([]terrain, error) {
	query := "select cellType from cell;"
	terrains := make([]terrain, 0)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var terrain terrain
		err := rows.Scan(&terrain)
		if err != nil {
			return nil, err
		}
		terrains = append(terrains, terrain)
	}
	fmt.Println("Terrain", terrains)
	return terrains, nil
}

func loadNations(db *sql.DB) ([]nation, error) {
	query := "select nationType from team;"
	nations := make([]nation , 0)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var nation nation
		err := rows.Scan(&nation)
		if err != nil {
			return nil, err
		}
		nations = append(nations, nation)
	}
	fmt.Println("Nations", nations)
	return nations, nil
}

func loadUnit(db *sql.DB, name string) (int, []*attack, *armor, *movement, error) {
	query := fmt.Sprintf("select health, attack_oneId, armorId, movementId from unit where name = \"%s\";", name)
	rows, err := db.Query(query)
	if err != nil {
		return 0, []*attack{}, nil, nil, err
	}
	var health int
	var attack_one_id int
	var armor_id int
	var movement_id int
	if rows.Next() {
		err := rows.Scan(&health, &attack_one_id, &armor_id, &movement_id)
		if err != nil {
			return 0, []*attack{}, nil, nil, err
		}
	}
	if rows.Next() {
		return 0, []*attack{}, nil, nil, errors.New("More than one entry returned in warrior query")
	}
	rows.Close()
	attack_one, err := loadAttack(db, attack_one_id)
	if err != nil {
		return 0, []*attack{}, nil, nil, err
	}
	armor, err := loadArmor(db, armor_id)
	if err != nil {
		return 0, []*attack{}, nil, nil, err
	}
	movement, err := loadMovement(db, movement_id)
	if err != nil {
		return 0, []*attack{}, nil, nil, err
	}
	return health, []*attack{attack_one}, armor, movement, nil
}

func loadAttack(db *sql.DB, attack_id int) (*attack, error) {
	query := fmt.Sprintf("select weaponTypeId, name, power, minRange, maxRange from weapon where id = %d;", attack_id)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var attackType_id int
	var name string
	var power int
	var minRange int
	var maxRange int
	if rows.Next() {
		err := rows.Scan(&attackType_id, &name, &power, &minRange, &maxRange)
		if err != nil {
			return nil, err
		}
	}
	if rows.Next() {
		return nil, errors.New("More than one entry returned in attack query")
	}
	rows.Close()
	attackType, err := loadAttackType(db, attackType_id)
	if err != nil {
		return nil, err
	}
	return newAttack(name, attackType, power, minRange, maxRange), nil
}

func loadAttackType(db *sql.DB, attackType_id int) (attackType, error) {
	query := fmt.Sprintf("select weaponType from weaponType where id = %d;", attackType_id)
	rows, err := db.Query(query)
	if err != nil {
		return 0, err
	}
	var dbAttackType int
	if rows.Next() {
		err := rows.Scan(&dbAttackType)
		if err != nil {
			return 0, err
		}
	}
	if rows.Next() {
		return 0, errors.New("More than one entry returned in attackType query")
	}
	rows.Close()
	return attackType(dbAttackType), nil
}

func loadArmor(db *sql.DB, armor_id int) (*armor, error) {
	query := fmt.Sprintf("select armorTypeId, name, strength from armor where id = %d;", armor_id)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var armorType_id int
	var name string
	var strength int
	if rows.Next() {
		err := rows.Scan(&armorType_id, &name, &strength)
		if err != nil {
			return nil, err
		}
	}
	if rows.Next() {
		return nil, errors.New("More than one entry returned in armor query")
	}
	rows.Close()
	armorType, err := loadArmorType(db, armorType_id)
	if err != nil {
		return nil, err
	}
	return newArmor(name, armorType, strength), nil
}

func loadArmorType(db *sql.DB, armorType_id int) (armorType, error) {
	query := fmt.Sprintf("select armorType from armorType where id = %d;", armorType_id)
	rows, err := db.Query(query)
	if err != nil {
		return 0, err
	}
	var dbArmorType int
	if rows.Next() {
		err := rows.Scan(&dbArmorType)
		if err != nil {
			return 0, err
		}
	}
	if rows.Next() {
		return 0, errors.New("More than one entry returned in armorType query")
	}
	rows.Close()
	return armorType(dbArmorType), nil
}

func loadMovement(db *sql.DB, movement_id int) (*movement, error) {
	query := fmt.Sprintf("select speedMapId, name, distance from movement where id = %d;", movement_id)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var speedMap_id int
	var name string
	var distance int
	if rows.Next() {
		err := rows.Scan(&speedMap_id, &name, &distance)
		if err != nil {
			return nil, err
		}
	}
	if rows.Next() {
		return nil, errors.New("More than one entry returned in movement query")
	}
	rows.Close()
	speedMap, err := loadSpeedMap(db, speedMap_id)
	if err != nil {
		return nil, err
	}
	return newMovement(name, distance, speedMap), nil
}

func loadSpeedMap(db *sql.DB, speedMap_id int) (map[terrain]multiplier, error) {
	query := fmt.Sprintf("select speeds from speedMap where id = %d;", speedMap_id)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var dbSpeedMap string
	if rows.Next() {
		err := rows.Scan(&dbSpeedMap)
		if err != nil {
			return nil, err
		}
	}
	if rows.Next() {
		return nil, errors.New("More than one entry returned in speedMap query")
	}
	rows.Close()
	var stringSpeedMap map[string]int
	fmt.Println("--------JSON-----------")
	fmt.Println(json.Unmarshal([]byte(dbSpeedMap), &stringSpeedMap))
	fmt.Println(stringSpeedMap)
	speedMap := make(map[terrain]multiplier, len(stringSpeedMap))
	for cellType, cost := range stringSpeedMap {
		intCell, err := strconv.Atoi(cellType)
		if err != nil {
			return nil, err
		}
		speedMap[terrain(intCell)] = multiplier(cost)
	}
	fmt.Println("--------JSON-----------")
	return speedMap, nil
}
