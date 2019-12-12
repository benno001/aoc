package main

import "testing"
import "reflect"

func TestParseInput(t *testing.T) {
	tables := []struct {
		input  []string
		output []moon
	}{
		{
			[]string{"<x=7, y=10, z=17>", "<x=-2, y=7, z=0>", "<x=12, y=5, z=12>", "<x=5, y=-8, z=6>"},
			[]moon{
				moon{position: point{x:7,y:10,z:17}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:-2,y:7,z:0}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:12,y:5,z:12}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:5,y:-8,z:6}, velocity: point{x:0,y:0,z:0}},
			},
		},
		{
			[]string{"<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>", "<x=4, y=-8, z=8>", "<x=3, y=5, z=-1>"},
			[]moon{
				moon{position: point{x:-1,y:0,z:2}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:2,y:-10,z:-7}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:4,y:-8,z:8}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:3,y:5,z:-1}, velocity: point{x:0,y:0,z:0}},
			},
		},
	}
	for _, table := range tables {
		result := parseInput(table.input)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.input, result, table.output)
		}
	}
}

func TestProcessSteps(t *testing.T) {
	tables := []struct {
		input  []moon
		steps int
		output []moon
	}{
		{
			[]moon{
				moon{position: point{x:-1,y:0,z:2}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:2,y:-10,z:-7}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:4,y:-8,z:8}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:3,y:5,z:-1}, velocity: point{x:0,y:0,z:0}},
			},
			0,
			[]moon{
				moon{position: point{x:-1,y:0,z:2}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:2,y:-10,z:-7}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:4,y:-8,z:8}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:3,y:5,z:-1}, velocity: point{x:0,y:0,z:0}},
			},
		},
		{
			[]moon{
				moon{position: point{x:-1,y:0,z:2}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:2,y:-10,z:-7}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:4,y:-8,z:8}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:3,y:5,z:-1}, velocity: point{x:0,y:0,z:0}},
			},
			1,
			[]moon{
				moon{position: point{x:2,y:-1,z:1}, velocity: point{x:3,y:-1,z:-1}},
				moon{position: point{x:3,y:-7,z:-4}, velocity: point{x:1,y:3,z:3}},
				moon{position: point{x:1,y:-7,z:5}, velocity: point{x:-3,y:1,z:-3}},
				moon{position: point{x:2,y:2,z:0}, velocity: point{x:-1,y:-3,z:1}},
			},
		},
		{
			[]moon{
				moon{position: point{x:-8,y:-10,z:0}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:5,y:5,z:10}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:2,y:-7,z:3}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:9,y:-8,z:-3}, velocity: point{x:0,y:0,z:0}},
			},
			10,
			[]moon{
				moon{position: point{x:-9,y:-10,z:1}, velocity: point{x:-2,y:-2,z:-1}},
				moon{position: point{x:4,y:10,z:9}, velocity: point{x:-3,y:7,z:-2}},
				moon{position: point{x:8,y:-10,z:-3}, velocity: point{x:5,y:-1,z:-2}},
				moon{position: point{x:5,y:-10,z:3}, velocity: point{x:0,y:-4,z:5}},
			},
		},
		{
			[]moon{
				moon{position: point{x:-8,y:-10,z:0}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:5,y:5,z:10}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:2,y:-7,z:3}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:9,y:-8,z:-3}, velocity: point{x:0,y:0,z:0}},
			},
			100,
			[]moon{
				moon{position: point{x:8,y:-12,z:-9}, velocity: point{x:-7,y:3,z:0}},
				moon{position: point{x:13,y:16,z:-3}, velocity: point{x:3,y:-11,z:-5}},
				moon{position: point{x:-29,y:-11,z:-1}, velocity: point{x:-3,y:7,z:4}},
				moon{position: point{x:16,y:-13,z:23}, velocity: point{x:7,y:1,z:1}},
			},
		},
	}
	for _, table := range tables {
		processSteps(table.input, table.steps)
		if !reflect.DeepEqual(table.input, table.output) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", table.input, table.output)
		}
	}
}

func TestCalculateSystemEnergy(t *testing.T){
	tables := []struct{
		moons []moon
		result int
	}{
		{
			[]moon{
				moon{position: point{x:8,y:-12,z:-9}, velocity: point{x:-7,y:3,z:0}},
				moon{position: point{x:13,y:-16,z:-3}, velocity: point{x:3,y:-11,z:-5}},
				moon{position: point{x:-29,y:-11,z:-1}, velocity: point{x:-3,y:7,z:4}},
				moon{position: point{x:16,y:-13,z:23}, velocity: point{x:7,y:1,z:1}},
			},
			1940,
		},
	}
	for _, table := range tables {
		result := calculateSystemEnergy(table.moons)
		if result != table.result {
			t.Errorf("Output was incorrect, got: %v, want: %v.", result, table.result)
		}
	}
}

func TestProcessGravity(t *testing.T) {
	tables := []struct{
		moons []moon
		result []moon
	}{
		{
			[]moon{
				moon{position: point{x:-1,y:0,z:2}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:2,y:-10,z:-7}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:4,y:-8,z:8}, velocity: point{x:0,y:0,z:0}},
				moon{position: point{x:3,y:5,z:-1}, velocity: point{x:0,y:0,z:0}},
			},
			[]moon{
				moon{position: point{x:-1,y:0,z:2}, velocity: point{x:3,y:-1,z:-1}},
				moon{position: point{x:2,y:-10,z:-7}, velocity: point{x:1,y:3,z:3}},
				moon{position: point{x:4,y:-8,z:8}, velocity: point{x:-3,y:1,z:-3}},
				moon{position: point{x:3,y:5,z:-1}, velocity: point{x:-1,y:-3,z:1}},
			},
		},
	}
	for _, table := range tables{
		processGravity(table.moons)
		if !reflect.DeepEqual(table.result, table.moons) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", table.moons, table.result)
		}
	}
}

func TestProcessVelocity(t *testing.T) {
	tables := []struct{
		moons []moon
		result []moon
	}{
		{
			[]moon{
				moon{position: point{x:-1,y:0,z:2}, velocity: point{x:3,y:-1,z:-1}},
				moon{position: point{x:2,y:-10,z:-7}, velocity: point{x:1,y:3,z:3}},
				moon{position: point{x:4,y:-8,z:8}, velocity: point{x:-3,y:1,z:-3}},
				moon{position: point{x:3,y:5,z:-1}, velocity: point{x:-1,y:-3,z:1}},
			},
			[]moon{
				moon{position: point{x:2,y:-1,z:1}, velocity: point{x:3,y:-1,z:-1}},
				moon{position: point{x:3,y:-7,z:-4}, velocity: point{x:1,y:3,z:3}},
				moon{position: point{x:1,y:-7,z:5}, velocity: point{x:-3,y:1,z:-3}},
				moon{position: point{x:2,y:2,z:0}, velocity: point{x:-1,y:-3,z:1}},
			},
		},
	}
	for _, table := range tables{
		processVelocity(table.moons)
		if !reflect.DeepEqual(table.result, table.moons) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", table.moons, table.result)
		}
	}
}

// func TestProcessVelocity(t *testing.T){

// }