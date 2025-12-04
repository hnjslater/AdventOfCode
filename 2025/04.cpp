

#include "array"
#include "fstream"
#include "ios"
#include "iostream"
#include "set"
#include "utility"
#include "vector"

int main(int argc, char** argv) {
  auto part = std::stoi(argv[1]);
  auto filename = argv[2];

  std::ifstream file(filename, std::ios_base::in);
  std::set<std::pair<int, int>> warehouse;
  std::string line;
  auto row = 0;
  while (getline(file, line)) {
    for (auto col = 0; col < line.length(); col++) {
      if (line[col] == '@') {
        warehouse.insert(std::make_pair(row, col));
      }
    }
    row++;
  }
  auto result = 0;

  std::set next_todo = warehouse;
  do {
    std::set todo = next_todo;
    next_todo.clear();
    for (auto const& coords : todo) {
      auto neighbours = 0;
      std::vector<std::pair<int, int>> candidates;
      for (auto& r : std::to_array<int>({-1, 0, 1})) {
        for (auto& c : std::to_array<int>({-1, 0, 1})) {
          auto neighbour = std::make_pair(coords.first + r, coords.second + c);
          if (warehouse.contains(neighbour)) {
            neighbours++;
            candidates.push_back(neighbour);
          }
        }
      }
      if (neighbours < 5) {
        result++;
        if (part == 2) {
          warehouse.erase(coords);
        }
        std::copy(candidates.begin(), candidates.end(),
                  std::inserter(next_todo, next_todo.end()));
      }
    }
  } while (!next_todo.empty() && part == 2);

  if (part == 2) {
    result /= 2;
  }
  std::cout << result << std::endl;
}
