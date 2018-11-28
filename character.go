package main

import (
	"math/rand"

	"github.com/ironarachne/utility"
)

type character struct {
	class        characterClass
	race         characterRace
	charisma     int
	constitution int
	dexterity    int
	intelligence int
	strength     int
	wisdom       int
	hitPoints    int
	spells       []spell
}

type characterClass struct {
	description string
	hitDie      int
	name        string
	isCaster    bool
}

type characterRace struct {
	adultAge    int
	description string
	maxAge      int
	name        string
}

type spell struct {
	name        string
	description string
	level       int
}

func randomCharacter() character {
	character := character{}

	character.charisma = rollDice(3, 6)
	character.constitution = rollDice(3, 6)
	character.dexterity = rollDice(3, 6)
	character.intelligence = rollDice(3, 6)
	character.strength = rollDice(3, 6)
	character.wisdom = rollDice(3, 6)

	character.class = randomClass()
	character.race = randomRace()

	if character.class.isCaster {

	}

	character.hitPoints = randomHitPoints(character.class.hitDie, character.constitution)

	return character
}

func randomClass() characterClass {
	className := utility.RandomItem(classNames)

	class := classes[className]

	return class
}

func randomHitPoints(hitDie int, constitution int) int {
	conBonus := 0

	if constitution > 12 {
		conBonus = 1
	}

	return rand.Intn(hitDie) + 1 + conBonus
}

func randomRace() characterRace {
	raceName := utility.RandomItem(raceNames)

	race := races[raceName]

	return race
}
