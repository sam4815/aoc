package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Resource int

const (
	Ore      Resource = 1
	Clay              = 2
	Obsidian          = 3
	Geode             = 4
)

type Robot struct {
	robotype Resource
}

type RoboCrew struct {
	robotypes map[Resource]int
	resources map[Resource]int
	minute    int
}

type Factory struct {
	id         int
	blueprints map[Resource](map[Resource]int)
	max_costs  map[Resource]int
}

func parseResource(resource string) Resource {
	switch resource {
	case "ore":
		return Ore
	case "clay":
		return Clay
	case "obsidian":
		return Obsidian
	case "geode":
		return Geode
	}
	return Ore
}

func parseBlueprint(blueprint string) Factory {
	blueprints := make(map[Resource]map[Resource]int)

	blueprint_id_str := strings.Split(strings.Split(blueprint, ": ")[0], " ")[1]
	blueprint_id, _ := strconv.Atoi(blueprint_id_str)

	robot_descriptions := strings.Split(strings.Split(blueprint, ": ")[1], ". ")

	for _, description := range robot_descriptions {
		robotype := parseResource(strings.Split(description, " ")[1])

		cost_description := strings.Split(description, "costs ")[1]
		costs := strings.Split(strings.Split(cost_description, ".")[0], " and ")

		cost_map := make(map[Resource]int)
		for _, cost := range costs {
			quantity_and_type := strings.Split(cost, " ")
			quantity, _ := strconv.Atoi(quantity_and_type[0])
			mineral_type := parseResource(quantity_and_type[1])
			cost_map[mineral_type] = quantity
		}

		blueprints[robotype] = cost_map
	}

	return Factory{blueprints: blueprints, id: blueprint_id}
}

func (factory *Factory) CalculateMaxCosts() {
	max_costs := make(map[Resource]int)
	for _, costs := range factory.blueprints {
		for resource, cost := range costs {
			if max_costs[resource] < cost {
				max_costs[resource] = cost
			}
		}
	}

	factory.max_costs = max_costs
}

func (robocrew RoboCrew) Copy() RoboCrew {
	robotypes := make(map[Resource]int)
	for robotype, quantity := range robocrew.robotypes {
		robotypes[robotype] = quantity
	}

	resources := make(map[Resource]int)
	for resource, quantity := range robocrew.resources {
		resources[resource] = quantity
	}

	return RoboCrew{minute: robocrew.minute, resources: resources, robotypes: robotypes}
}

func (robocrew *RoboCrew) Tick() {
	robocrew.minute += 1

	for robotype, quantity := range robocrew.robotypes {
		robocrew.resources[robotype] += quantity
	}
}

func printResource(resource Resource) string {
	switch resource {
	case Ore:
		return "ore"
	case Clay:
		return "clay"
	case Obsidian:
		return "obsidian"
	case Geode:
		return "geode"
	}
	return ""
}

func (robocrew RoboCrew) Print() {
	minute_string := fmt.Sprintf("Minute: %d. ", robocrew.minute)
	robo_string := "Robots: "
	resource_string := "Resources: "

	for robotype, quantity := range robocrew.robotypes {
		robo_string += fmt.Sprintf("%d %s. ", quantity, printResource(robotype))
	}
	for resource, quantity := range robocrew.resources {
		resource_string += fmt.Sprintf("%d %s. ", quantity, printResource(resource))
	}

	log.Print(minute_string, robo_string, resource_string)
}

func (robocrew RoboCrew) Stringify() string {
	robo_string := ""
	resource_string := ""

	for robotype, quantity := range robocrew.robotypes {
		robo_string += fmt.Sprintf("t:%d,%d.", quantity, robotype)
	}
	for resource, quantity := range robocrew.resources {
		resource_string += fmt.Sprintf("t:%d,%d.", quantity, resource)
	}

	return robo_string + resource_string
}

func (robocrew *RoboCrew) Spend(cost map[Resource]int) {
	for resource, quantity := range cost {
		robocrew.resources[resource] -= quantity
	}
}

func (robocrew *RoboCrew) AddRobotype(resource Resource) {
	robocrew.robotypes[resource] += 1
}

func (robocrew RoboCrew) CanAfford(cost map[Resource]int) bool {
	can_afford := true

	for resource, quantity := range cost {
		if robocrew.resources[resource] < quantity {
			can_afford = false
			break
		}
	}

	return can_afford
}

// Various hacky ways to discard paths
func (robocrew RoboCrew) SpendingUnwisely(factory Factory) bool {
	if robocrew.resources[Clay] > factory.max_costs[Clay]*3 && robocrew.resources[Ore] > factory.max_costs[Ore]*3 {
		return true
	}

	if robocrew.resources[Ore] > factory.blueprints[Geode][Ore]*2 && robocrew.resources[Obsidian] > factory.blueprints[Geode][Obsidian]*2 {
		return true
	}

	if robocrew.resources[Ore] > factory.max_costs[Ore]*3 {
		return true
	}

	if robocrew.robotypes[Ore] > factory.max_costs[Ore] {
		return true
	}

	if robocrew.robotypes[Clay] > factory.max_costs[Clay] {
		return true
	}

	if robocrew.robotypes[Obsidian] > factory.max_costs[Obsidian] {
		return true
	}

	if robocrew.minute > 16 {
		ratio := factory.blueprints[Obsidian][Clay] / factory.blueprints[Obsidian][Ore]
		if ratio >= 3 && robocrew.robotypes[Clay]/robocrew.robotypes[Ore] < 1 {
			return true
		}

		if ratio <= 1 && robocrew.robotypes[Clay]/robocrew.robotypes[Ore] >= 2 {
			return true
		}
	}

	return false
}

func (robocrew *RoboCrew) GetChoices(factory Factory) []RoboCrew {
	choices := make([]RoboCrew, 0)

	for i := 4; i >= 1; i-- {
		if robocrew.CanAfford(factory.blueprints[Resource(i)]) {
			copy := robocrew.Copy()

			copy.Spend(factory.blueprints[Resource(i)])
			copy.Tick()
			copy.AddRobotype(Resource(i))

			if copy.SpendingUnwisely(factory) {
				continue
			}

			choices = append(choices, copy)
		}
	}

	// Do nothing
	copy := robocrew.Copy()
	copy.Tick()

	if copy.SpendingUnwisely(factory) {
		return choices
	}

	return append(choices, copy)
}

func (robocrew RoboCrew) FindMaxGeodes(factory Factory) int {
	robocrew_queue, curr_robocrew := []RoboCrew{robocrew}, robocrew

	max_geodes := 0
	best_obsidian_by_minute_map := make(map[int]int)
	best_geode_by_minute_map := make(map[int]int)
	traversal_map := make(map[string]bool)

	for len(robocrew_queue) > 0 {
		curr_robocrew, robocrew_queue = robocrew_queue[0], robocrew_queue[1:]

		if curr_robocrew.minute == 32 {
			if curr_robocrew.resources[Geode] > max_geodes {
				max_geodes = curr_robocrew.resources[Geode]
			}

			continue
		}

		robocrew_id := curr_robocrew.Stringify()
		if traversal_map[robocrew_id] == true {
			continue
		} else {
			traversal_map[robocrew_id] = true
		}

		if curr_robocrew.robotypes[Obsidian] < best_obsidian_by_minute_map[curr_robocrew.minute] &&
			curr_robocrew.robotypes[Geode] < best_geode_by_minute_map[curr_robocrew.minute] {
			continue
		} else {
			if curr_robocrew.robotypes[Obsidian] > best_obsidian_by_minute_map[curr_robocrew.minute] {
				best_obsidian_by_minute_map[curr_robocrew.minute] = curr_robocrew.robotypes[Obsidian]
			}
			if curr_robocrew.robotypes[Geode] > best_geode_by_minute_map[curr_robocrew.minute] {
				best_geode_by_minute_map[curr_robocrew.minute] = curr_robocrew.robotypes[Geode]
			}
		}

		possible_paths := curr_robocrew.GetChoices(factory)

		robocrew_queue = append(possible_paths, robocrew_queue...)
	}

	return max_geodes
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	quality_level := 0
	geode_product := 1
	factories := make([]Factory, 0)

	for scanner.Scan() {
		factory := parseBlueprint(scanner.Text())
		factory.CalculateMaxCosts()
		factories = append(factories, factory)
	}

	for _, factory := range factories {
		one_ore_robot := map[Resource]int{Ore: 1}
		robocrew := RoboCrew{robotypes: one_ore_robot, minute: 0, resources: make(map[Resource]int)}
		max_geodes := robocrew.FindMaxGeodes(factory)
		log.Print("Processed blueprint ", factory.id, ", found ", max_geodes, " geodes")
		quality_level += max_geodes * factory.id
		geode_product *= max_geodes
	}

	time_elapsed := time.Since(start)

	log.Printf(`
The sum of the quality levels is %d.
The product of the maximum number of geodes is %d.
Solution generated in %s.`,
		quality_level,
		geode_product,
		time_elapsed,
	)
}
