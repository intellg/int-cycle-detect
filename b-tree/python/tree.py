class Tree:
    def __init__(self, m=3):
        self.root = None

        if m < 3:
            self.m = 3
        else:
            self.m = m

        self.underflow = int((m - 1) / 2)  # int((m-1)/2) == ceil(m/2)-1

    def search_key(self, key, node=None, deep=0):
        deep += 1
        if not node:
            node = self.root

        search_result = node.search_key(key)
        if search_result['is_found']:
            return {'is_found': True,
                    'node': node,
                    'index': search_result['index'],
                    'deep': deep}
        else:
            if not node.children:  # node is leaf
                return {'is_found': False,
                        'node': node,
                        'index': search_result['index'],
                        'deep': deep}
            else:
                return self.search_key(key, node.children[search_result['index']], deep)

    def add_key(self, key):
        if self.root:
            search_result = self.search_key(key)
            if search_result['is_found']:
                return  # The key already existed. No further action.
            else:
                search_result['node'].add_key(self, key, search_result['index'] + 1)
        else:
            self.root = Node(key)

    def delete_key(self, key):
        if self.root:
            search_result = self.search_key(key)
            if search_result['is_found']:
                search_result['node'].delete_key(self, search_result['index'])
            else:
                return  # The key doesn't exist. No further action.

    def validate(self, node, step=0, deep=None):
        step += 1

        # The key_list[0] stores exact the count of the keys
        if not node.key_list[0] == len(node.key_list) - 1:
            return False

        # The count of keys is smaller than m and bigger than ceil(m/2)-1
        if node.parent is None:  # root node
            if not node.children:  # root & leaf
                if not node.key_list[0] >= 0:
                    return False
            else:
                if not 0 < node.key_list[0] < self.m:
                    return False
        else:  # non-root node
            if not self.underflow <= node.key_list[0] < self.m:
                return False

        # All leaf nodes have the same deep
        if not node.children:  # node is leaf
            if deep is None:
                deep = step
            else:
                if not step == deep:
                    return False

        # The keys in key_list is ascending
        for i in range(1, node.key_list[0]):
            if not node.key_list[i] < node.key_list[i + 1]:
                return False

        if node.children:
            for i in range(0, node.key_list[0] + 1):
                # The index field stores the correct value
                if not node.children[i].index == i:
                    return False
                # The parent refers to the correct object
                if not node.children[i].parent is node:
                    return False

            for i in range(1, node.key_list[0]):
                # Each key is bigger than the left child's biggest key and smaller than the right child's smallest key
                if not node.children[i - 1].key_list[-1] < node.key_list[i] < node.children[i].key_list[1]:
                    return False

            # Traverse all nodes
            for child in node.children:
                return self.validate(child, step, deep)

        step -= 1
        return True


class Node:
    def __init__(self, key=None, key_list=None, children=None):
        self.parent = None  # Node
        self.index = -1  # int

        if key:
            self.key_list = [1, key]
        elif key_list:
            self.key_list = key_list
            self.key_list.insert(0, len(key_list))

        self.children = children  # Node[]
        if children:
            for i in range(0, self.key_list[0] + 1):
                children[i].index = i

    def get_leftest_leaf_from(self, index):
        node = self.children[index]
        while node.children:
            node = node.children[0]
        return node

    def search_key(self, key):
        end = len(self.key_list)
        if end <= 1:
            return {'is_found': False,
                    'index': 1,
                    'step': 0}
        start = 1

        previous_mid = None
        step = 0
        while True:
            step += 1

            mid = (start + end) / 2
            int_mid = int(mid)
            if mid == previous_mid:
                if mid == int_mid:
                    int_mid -= 1
                return {'is_found': False,
                        'index': int_mid,
                        'step': step}
            else:
                previous_mid = mid

            if key == self.key_list[int_mid]:
                return {'is_found': True,
                        'index': int_mid,
                        'step': step}
            elif key < self.key_list[int_mid]:
                end = int_mid
            else:  # if key > self.key_list[int_mid]:
                start = int_mid

    def add_key(self, tree, key, index, child=None):
        self.key_list.insert(index, key)
        self.key_list[0] += 1

        if self.children:
            for i in range(index, len(self.children)):
                self.children[i].index += 1

            self.children.insert(index, child)  # child should never be None here
            child.parent = self
            child.index = index

        # Separation
        if self.key_list[0] == tree.m:
            return self.__adjust_overflow(tree)

    def __adjust_overflow(self, tree):
        # 0. Prepare necessary data
        mid = int(tree.m / 2) + 1
        new_node = Node(key_list=self.key_list[mid + 1:], children=self.children[mid:] if self.children else None)

        # 1. Handle parent node
        if self.parent:
            # Node self.key_list[mid] is upgraded
            self.parent.add_key(tree, self.key_list[mid], self.index + 1, new_node)
        else:
            tree.root = Node(key=self.key_list[mid], children=[self, new_node])
            self.parent = tree.root
            new_node.parent = tree.root

        # 2. Handle new node
        if new_node.children:
            for child in new_node.children:
                child.parent = new_node

        # 3. Handle self node
        self.key_list = self.key_list[0:mid]
        self.key_list[0] = mid - 1
        if self.children:
            self.children = self.children[0:mid]

    def delete_key(self, tree, index):
        if self.children:
            # Change the key with the right child's smallest key
            right_child_smallest_leaf = self.get_leftest_leaf_from(index)
            self.key_list[index] = right_child_smallest_leaf.key_list[1]
            delete_node = right_child_smallest_leaf
            delete_index = 1
        else:  # self node is leaf
            delete_node = self
            delete_index = index

        # Remove the key from the leaf node then adjust underflow if any
        delete_node.key_list.pop(delete_index)
        delete_node.key_list[0] -= 1
        if delete_node.key_list[0] < tree.underflow:
            delete_node.__adjust_underflow(tree)

    def __adjust_underflow(self, tree):
        if self.parent is None:
            if self.key_list[0] == 0:
                if self.children:
                    tree.root = self.children[0]
                    tree.root.index = -1
                    tree.root.parent = None
            return

        # 0. Borrow from brother logic
        if self.index > 0:  # Try to borrow from left
            left_node = self.parent.children[self.index - 1]
            if left_node.key_list[0] > tree.underflow:
                self.__borrow_from_brother(left_node, 1, self.index, -1)
                return
            right_node = self
        else:  # Try to borrow from right
            right_node = self.parent.children[self.index + 1]
            if right_node.key_list[0] > tree.underflow:
                self.__borrow_from_brother(right_node, self.key_list[0] + 1, self.index + 1, 1)
                return
            left_node = self

        # 1. Combine right node into left node
        # 1.1 Handle left node's children
        if left_node.children:
            for child in right_node.children:
                left_node.children.append(child)
                child.parent = left_node
                child.index += left_node.key_list[0] + 1

        # 1.2 Handle left node's key_list
        left_node.key_list.append(left_node.parent.key_list[right_node.index])
        left_node.key_list.extend(right_node.key_list[1:])
        left_node.key_list[0] += right_node.key_list[0] + 1

        # 2. Handle parent node
        # 2.1 Handle parent node's children
        left_node.parent.children.pop(right_node.index)
        if left_node.parent.key_list[0] > left_node.index:
            for child in left_node.parent.children[right_node.index:]:
                child.index -= 1

        # 2.2 Handle parent node's key_list
        left_node.parent.key_list.pop(right_node.index)
        left_node.parent.key_list[0] -= 1
        if left_node.parent.key_list[0] < tree.underflow:
            left_node.parent.__adjust_underflow(tree)

    def __borrow_from_brother(self, brother, self_index, parent_index, brother_index):
        # 1. Move the key from parent node to self node
        self.key_list.insert(self_index, self.parent.key_list[parent_index])
        self.key_list[0] += 1

        # 2. Move the key from brother node to parent node
        self.parent.key_list[parent_index] = brother.key_list.pop(brother_index)
        brother.key_list[0] -= 1

        # 3. Move the child from brother to current node
        if self.children:
            if brother_index == -1:  # Borrow from left
                self.children.insert(0, brother.children.pop(-1))
                self.children[0].parent = self
                for i in range(0, self.key_list[0] + 1):
                    self.children[i].index = i
            else:  # Borrow from right
                self.children.append(brother.children.pop(0))
                self.children[-1].parent = self
                self.children[-1].index = self.key_list[0]
                for i in range(0, brother.key_list[0] + 1):
                    brother.children[i].index = i


def main():
    tree = Tree(3)

    tree.add_key(10)
    tree.add_key(20)
    tree.add_key(30)
    tree.add_key(40)
    tree.add_key(50)
    tree.add_key(32)
    tree.add_key(35)
    tree.add_key(55)
    tree.add_key(60)

    tree.validate(tree.root)
    print(tree)


if __name__ == '__main__':
    main()
