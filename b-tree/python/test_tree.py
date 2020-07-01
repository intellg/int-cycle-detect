import random
import unittest

from tree import Tree


class TestTree(unittest.TestCase):
    test_round = 0

    # Testing meta data
    max_m = 20
    test_count = 50
    test_range = 1000

    def test_tree(self):
        print('# ======== Start testing test_tree() ========')
        for m in range(3, TestTree.max_m + 1):
            tree = Tree(m)
            node_set = set()

            # Construct b-tree
            while len(node_set) < TestTree.test_count:
                key = random.randint(1, TestTree.test_range)
                # Below line is only used for debugging
                # print("    tree.add_key({})".format(key))

                tree.add_key(key)
                node_set.add(key)

                # Test b-tree
                self.perform_testing(tree, node_set)

            # Un-construct b-tree
            while len(node_set) > 0:
                key = node_set.pop()
                # Below line is only used for debugging
                # print("    tree.delete_key({})".format(key))

                tree.delete_key(key)

                # Test b-tree
                self.perform_testing(tree, node_set)

    def perform_testing(self, tree, node_set):
        TestTree.test_round += 1
        print("# ======== test round {} with m={} ========".format(TestTree.test_round, tree.m))
        self.assertTrue(tree.validate(tree.root))
        self.check_keys(tree, node_set, TestTree.test_range)

    def check_keys(self, tree, node_set, test_range):
        for i in range(1, test_range + 1):
            search_result = tree.search_key(i)
            if i in node_set:
                self.assertTrue(search_result['is_found'])
            else:
                self.assertFalse(search_result['is_found'])


if __name__ == '__main__':
    unittest.main()
