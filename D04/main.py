from dataclasses import dataclass

@dataclass
class Coord:
  x: int
  y: int

  def __add__(self, o):
    return Coord(self.x + o.x, self.y + o.y)

  def in_bounds(self, bounds):
    return self.x >= 0 and self.x < bounds.x and self.y >= 0 and self.y < bounds.y

def xmas_search(start: Coord, dir: Coord, lines: list[str]) -> bool:
  bounds = Coord(len(lines), len(lines[0]))
  m = start + dir
  a = m + dir
  s = a + dir
  if not (m.in_bounds(bounds) and a.in_bounds(bounds) and s.in_bounds(bounds)):
    return False
  return lines[m.x][m.y] == 'M' and lines[a.x][a.y] == 'A' and lines[s.x][s.y] == 'S'

DIRS = [
    Coord(1, 0),
    Coord(-1, 0),
    Coord(0, 1),
    Coord(0, -1),
    Coord(1, 1),
    Coord(-1, 1),
    Coord(1, -1),
    Coord(-1, -1)
]

def part_1(filename='./inputs/day_4.txt'):
  with open(filename) as file:
    lines = [line.strip() for line in file.readlines()]
    xmas_count = 0
    for row, line in enumerate(lines):
      for col, c in enumerate(line):
        if c == 'X':
          for dir in DIRS:
            xmas_count += xmas_search(Coord(row, col), dir, lines)

    return xmas_count

print(part_1('input')) # 18