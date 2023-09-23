#include <iostream>
#include <vector>

bool sumToTargetRec(const std::vector<int> &aVec, const int target, int depth) {
	if (target == 0) return true;
    if (depth >= aVec.size() || target < 0) return false;
        
    for (int i { depth }; i < aVec.size(); ++i) {
        const int newTarget { target - aVec.at(depth) };
        
        if (sumToTargetRec(aVec, newTarget, i + 1)) return true;
    }
	
	return false;
}

bool sumToTarget(const std::vector<int> &aVec, const int target) {
	return sumToTargetRec(aVec, target, 0);
}

int main (void) {
	std::vector<int> aVec { 1, 3, 7 };
	int kTarget { 8 };

	if (sumToTarget(aVec, kTarget))
		std::cout << "yes" << std::endl;
	else
		std::cout << "no" << std::endl;

	kTarget = 6;

	if (sumToTarget(aVec, kTarget))
		std::cout << "yes" << std::endl;
	else
		std::cout << "no" << std::endl;

	return 0;
}
