package entities

import (
  "logger"
  "db"
)

type Command struct {
  Name string
  ArgCount int
  ArgRegex string
}

type Item struct {
  Id int
  Name string
  Description string
}

type Player struct {
  Id int
  Name string
  Dex int
  Str int
  Int int
  Wis int
  Con int
  Chr int
  Room int
}

type Inventory struct {
  Id int
  Player int
  Item int
  Quantity int
}

type Room struct {
  Id int
  Name string
  Description string
}

type Transition struct {
  Id int
  Source int
  Target int
  Command string
  CommandArgs string
}

func SetupTables() {
  var table db.TableDefinition
  var index map[string][]int64

  // Command
  logger.Info("Creating commands table")
  db.CreateTableIfNotExist("commands", []string{
    "Name",
    "ArgCount",
    "ArgRegex",
  }, []string{
    "string",
    "integer",
    "string",
  }, 0, true)

  // Item
  logger.Info("Creating items table")
  db.CreateTableIfNotExist("items", []string{
    "Id",
    "Name",
    "Description",
  }, []string{
    "integer",
    "string",
    "string",
  }, 0, true)

  // Player
  logger.Info("Creating players table")
  db.CreateTableIfNotExist("players", []string{
    "Id",
    "Name",
    "Dex",
    "Str",
    "Int",
    "Wis",
    "Con",
    "Chr",
    "Room",
  }, []string{
    "integer",
    "string",
    "integer",
    "integer",
    "integer",
    "integer",
    "integer",
    "integer",
    "integer",
  }, 1, true)

  // Inventoriy
  logger.Info("Creating inventory table")
  table = db.CreateTableIfNotExist("inventory", []string{
    "Id",
    "Player",
    "Item",
    "Quantity",
  }, []string{
    "integer",
    "integer",
    "integer",
    "integer",
  }, 0, false)
  index = db.CreateIndex(table.CSV, "Item")
  table.Info.Indices["Item"] = index

  // Room
  logger.Info("Creating rooms table")
  db.CreateTableIfNotExist("rooms", []string{
    "Id",
    "Name",
    "Description",
  }, []string{
    "integer",
    "string",
    "string",
  }, 0, true)

  // Transition
  table = db.CreateTableIfNotExist("transitions", []string{
    "Id",
    "Source",
    "Target",
    "Command",
    "CommandArgs",
  }, []string{
    "integer",
    "integer",
    "integer",
    "string",
    "string",
  }, 0, true)
  index = db.CreateIndex(table.CSV, "Target")
  table.Info.Indices["Target"] = index
  index = db.CreateIndex(table.CSV, "Source")
  table.Info.Indices["Source"] = index
}
