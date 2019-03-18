/*
在一个二维数组中（每个一维数组的长度相同），
每一行都按照从左到右递增的顺序排序，
每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/

#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

/*
每一行和每一列都是有序的, 并且从左至右, 从上至下都是递增

都可以通过判断大小过滤一行或者一列
*/

class Solution {
public:

    // 暴力解决
    bool Find(int target, vector<vector<int> > array) {

        int i, j;
        for (i = 0;i < array.size(); ++i) {
            for (j = 0;j < array[i].size(); ++j) {
                if (array[i][j] == target) {
                    return true;
                }
                printf("%d\t\t", array[i][j]);
            }
            printf("\n");
        }

        return false;
    }

    // 从右上角开始
    bool FindRighttop(int target, vector<vector<int> > array) {

        int i = 0,j = 0;

        // 初始化 j
        if (array.size() > 0) {
            j = array[0].size() - 1;
        }else{
            return false;
        }

        for (i = 0;i < array.size() && j >= 0; ) {
            if (array[i][j] == target) {
                return true;
            }else if (array[i][j] < target) {
                ++i;
            }else {
                --j;
            }
        }
        return false;
    }

    // 从左下角开始
    bool FindLefttop(int target, vector<vector<int> > array) {

        int i = 0,j = 0;

        for (i = array.size() - 1 ;i >= 0 && j < array[i].size(); ) {
            if (array[i][j] == target) {
                return true;
            }else if (array[i][j] > target) {
                --i;
            }else {
                ++j;
            }
        }
        return false;
    }

    void Print(vector<vector<int> > array) {
         int i, j;
         for (i = 0;i < array.size(); ++i) {
            for (j = 0;j < array[i].size(); ++j) {
                printf("%d\t\t", array[i][j]);
            }
            printf("\n");
        }
    }

};

int main() {

    Solution a;

    vector<vector<int> > array;
    int count = 0;
    for (int i = 0; i < 4; ++i) {
        vector<int> a;
        for (int j = 0; j < 4; ++j) {
            a.push_back(++count);
        }
        array.push_back(a);
    }
    a.Print(array);
    std::cout << a.FindRighttop(21, array) << std::endl;
    std::cout << a.FindLefttop(21, array) << std::endl;

    return 0;
}