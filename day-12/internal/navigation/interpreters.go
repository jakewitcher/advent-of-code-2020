package navigation

type Interpreter interface {
	RunInstruction(ship *Ship, instruction *Instruction)
}

type NaiveInterpreter struct{}

func (n NaiveInterpreter) RunInstruction(s *Ship, instruction *Instruction) {
	switch instruction.Action {
	case 'N':
		s.Y += instruction.Units
	case 'S':
		s.Y -= instruction.Units
	case 'E':
		s.X += instruction.Units
	case 'W':
		s.X -= instruction.Units
	case 'L':
		n.rotateLeft(s, instruction.Units)
	case 'R':
		n.rotateRight(s, instruction.Units)
	case 'F':
		n.RunInstruction(s, &Instruction{Action: s.Facing, Units: instruction.Units})
	}
}

func (NaiveInterpreter) rotateLeft(s *Ship, degrees int) {
	directions := []rune{'N', 'W', 'S', 'E'}
	turns := degrees / 90
	var current int
	for i, direction := range directions {
		if direction == s.Facing {
			current = i
			break
		}
	}

	next := (current + turns) % 4
	s.Facing = directions[next]
}

func (NaiveInterpreter) rotateRight(s *Ship, degrees int) {
	directions := []rune{'N', 'E', 'S', 'W'}
	turns := degrees / 90
	var current int
	for i, direction := range directions {
		if direction == s.Facing {
			current = i
			break
		}
	}

	next := (current + turns) % 4
	s.Facing = directions[next]
}

type WaypointInterpreter struct{}

func (w WaypointInterpreter) RunInstruction(s *Ship, instruction *Instruction) {
	switch instruction.Action {
	case 'N':
		s.Waypoint.Y += instruction.Units
	case 'S':
		s.Waypoint.Y -= instruction.Units
	case 'E':
		s.Waypoint.X += instruction.Units
	case 'W':
		s.Waypoint.X -= instruction.Units
	case 'L':
		w.rotateLeft(s, instruction.Units)
	case 'R':
		w.rotateRight(s, instruction.Units)
	case 'F':
		w.moveShipTowardWaypoint(s, instruction.Units)
	}
}

func (WaypointInterpreter) rotateLeft(s *Ship, degrees int) {
	turns := degrees / 90

	for i := 0; i < turns; i++ {
		s.Waypoint.X, s.Waypoint.Y = -s.Waypoint.Y, s.Waypoint.X
	}
}

func (WaypointInterpreter) rotateRight(s *Ship, degrees int) {
	turns := degrees / 90
	for i := 0; i < turns; i++ {
		s.Waypoint.X, s.Waypoint.Y = s.Waypoint.Y, -s.Waypoint.X
	}
}

func (WaypointInterpreter) moveShipTowardWaypoint(s *Ship, units int) {
	s.X += units * s.Waypoint.X
	s.Y += units * s.Waypoint.Y
}
