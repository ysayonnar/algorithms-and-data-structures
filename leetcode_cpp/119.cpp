#include <iostream>
#include <vector>

void printVector(std::vector<int> vec) {
    std::cout << "Vector:" << std::endl;
    for (int i = 0; i < vec.size(); ++i) {
        std::cout << vec[i] << " ";
    }
    std::cout << std::endl;
}

std::vector<int> getRow(int rowIdx) {
    std::vector<int> current = {1, 1};
    std::vector<int> prev = {1};

    if (rowIdx == 0) {
        return prev;
    } else if (rowIdx == 1) {
        return current;
    }

    for (int i = 1; i < rowIdx; ++i) {
        prev.clear();
        for (int j = 0; j < current.size(); ++j) {
            prev.push_back(current[j]);
        }

        current.clear();
        current.resize(prev.size() + 1);
        current[0] = 1;
        current[prev.size()] = 1;

        for (int j = 0; j < prev.size() - 1; ++j) {
            current[j + 1] = prev[j] + prev[j + 1];
        }
    }

    return current;
}

int main() {
    int rowIdx;

    std::cout << "Inser row index:\t";
    std::cin >> rowIdx;

    std::vector<int> vec = getRow(rowIdx);
    printVector(vec);
}
