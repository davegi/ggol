package ggol

import (
	"fmt"
	"testing"
)

func testAreHasLiveCellTestUnitsMapsEqualCaseOne(t *testing.T) {
	g1 := unitsHavingLiveCellForTest{{true, false}, {true, false}}
	g2 := unitsHavingLiveCellForTest{{true, false}, {true, false}}

	if areTwoUnitsHavingLiveCellForTestEqual(g1, g2) {
		t.Log("Passed")
	} else {
		t.Fatalf("g1 and g2 should be equal.")
	}
}

func testAreHasLiveCellTestUnitsMapsEqualCaseTwo(t *testing.T) {
	g1 := unitsHavingLiveCellForTest{{true, false}, {true, false}}
	g2 := unitsHavingLiveCellForTest{{true, false}, {true, true}}

	if !areTwoUnitsHavingLiveCellForTestEqual(g1, g2) {
		t.Log("Passed")
	} else {
		t.Fatalf("g1 and g2 should not be equal.")
	}
}

func TestHasLiveCellTestUnitsMapsEqual(t *testing.T) {
	testAreHasLiveCellTestUnitsMapsEqualCaseOne(t)
	testAreHasLiveCellTestUnitsMapsEqualCaseTwo(t)
}

func testConvertTestUnitsMatrixToHasLiveCellTestUnitsMapCaseOne(t *testing.T) {
	uniMatrix := generateInitialUnitMatrixForTest(2, 2, unitForTest{hasLiveCell: true})
	game, _ := NewGame(uniMatrix)
	game.SetNextUnitGenerator(defaultUnitForTestIterator)
	allUnits := game.GetUnits()
	fmt.Println(allUnits)
	liveUnitsMap := convertUnitForTestMatrixToUnitsHavingLiveCellForTest(allUnits)

	expectedMap := unitsHavingLiveCellForTest{{true, true}, {true, true}}

	if areTwoUnitsHavingLiveCellForTestEqual(*liveUnitsMap, expectedMap) {
		t.Log("Passed")
	} else {
		t.Fatalf("Did not convert matrix of *TestUnit to unitsHavingLiveCellForTest successfully.")
	}
}

func TestConvertTestUnitsMatrixToHasLiveCellTestUnitsMap(t *testing.T) {
	testConvertTestUnitsMatrixToHasLiveCellTestUnitsMapCaseOne(t)
}
