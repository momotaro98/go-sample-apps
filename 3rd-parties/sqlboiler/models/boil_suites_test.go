// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Occupations", testOccupations)
	t.Run("Positions", testPositions)
	t.Run("Userlocations", testUserlocations)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Occupations", testOccupationsDelete)
	t.Run("Positions", testPositionsDelete)
	t.Run("Userlocations", testUserlocationsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Occupations", testOccupationsQueryDeleteAll)
	t.Run("Positions", testPositionsQueryDeleteAll)
	t.Run("Userlocations", testUserlocationsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Occupations", testOccupationsSliceDeleteAll)
	t.Run("Positions", testPositionsSliceDeleteAll)
	t.Run("Userlocations", testUserlocationsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Occupations", testOccupationsExists)
	t.Run("Positions", testPositionsExists)
	t.Run("Userlocations", testUserlocationsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Occupations", testOccupationsFind)
	t.Run("Positions", testPositionsFind)
	t.Run("Userlocations", testUserlocationsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Occupations", testOccupationsBind)
	t.Run("Positions", testPositionsBind)
	t.Run("Userlocations", testUserlocationsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Occupations", testOccupationsOne)
	t.Run("Positions", testPositionsOne)
	t.Run("Userlocations", testUserlocationsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Occupations", testOccupationsAll)
	t.Run("Positions", testPositionsAll)
	t.Run("Userlocations", testUserlocationsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Occupations", testOccupationsCount)
	t.Run("Positions", testPositionsCount)
	t.Run("Userlocations", testUserlocationsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Occupations", testOccupationsHooks)
	t.Run("Positions", testPositionsHooks)
	t.Run("Userlocations", testUserlocationsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Occupations", testOccupationsInsert)
	t.Run("Occupations", testOccupationsInsertWhitelist)
	t.Run("Positions", testPositionsInsert)
	t.Run("Positions", testPositionsInsertWhitelist)
	t.Run("Userlocations", testUserlocationsInsert)
	t.Run("Userlocations", testUserlocationsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("UserlocationToUserUsingUserId", testUserlocationToOneUserUsingUserId)
	t.Run("UserToPositionUsingPositionId", testUserToOnePositionUsingPositionId)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {
	t.Run("UserToUserlocationUsingUserIdUserlocation", testUserOneToOneUserlocationUsingUserIdUserlocation)
}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("PositionToPositionIdUsers", testPositionToManyPositionIdUsers)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("UserlocationToUserUsingUserIdUserlocation", testUserlocationToOneSetOpUserUsingUserId)
	t.Run("UserToPositionUsingPositionIdUsers", testUserToOneSetOpPositionUsingPositionId)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("UserToPositionUsingPositionIdUsers", testUserToOneRemoveOpPositionUsingPositionId)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {
	t.Run("UserToUserlocationUsingUserIdUserlocation", testUserOneToOneSetOpUserlocationUsingUserIdUserlocation)
}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("PositionToPositionIdUsers", testPositionToManyAddOpPositionIdUsers)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("PositionToPositionIdUsers", testPositionToManySetOpPositionIdUsers)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("PositionToPositionIdUsers", testPositionToManyRemoveOpPositionIdUsers)
}

func TestReload(t *testing.T) {
	t.Run("Occupations", testOccupationsReload)
	t.Run("Positions", testPositionsReload)
	t.Run("Userlocations", testUserlocationsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Occupations", testOccupationsReloadAll)
	t.Run("Positions", testPositionsReloadAll)
	t.Run("Userlocations", testUserlocationsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Occupations", testOccupationsSelect)
	t.Run("Positions", testPositionsSelect)
	t.Run("Userlocations", testUserlocationsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Occupations", testOccupationsUpdate)
	t.Run("Positions", testPositionsUpdate)
	t.Run("Userlocations", testUserlocationsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Occupations", testOccupationsSliceUpdateAll)
	t.Run("Positions", testPositionsSliceUpdateAll)
	t.Run("Userlocations", testUserlocationsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}