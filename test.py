import unittest
import entities
import module2migrate as m2m

class TestModule2Migrate(unittest.TestCase):

    def test_create_object(self):
        self.assertEqual(entities.TestObj(1, 'a'), m2m.create_object(1, 'a'))
        self.assertNotEqual(entities.TestObj(1, 'a'), m2m.create_object(2, 'a'))

    def test_num_times_string(self):
        obj = entities.TestObj(10, 'a')
        self.assertEqual('a' * 10, m2m.num_times_string(obj))

    def test_clone(self):
        obj = entities.TestObj(10, 'a')
        self.assertEqual(obj, m2m.clone(obj))
        self.assertNotEqual(obj, m2m.clone(obj, 2))

    def test_list_of_clones(self):
        obj = entities.TestObj(10, 'a')
        self.assertEqual([obj, obj, obj], m2m.list_of_clones(obj, 3))

    def test_fibinacci(self):
        obj = entities.TestObj(10, 'a')
        self.assertEqual(34, m2m.fibonacchi(obj))

    def test_md5(self):
        obj = entities.TestObj(10, 'abcdef')
        self.assertEqual('e80b5017098950fc58aad83c8c14978e', m2m.md5(obj))

    def test_catch_explode(self):
        obj = entities.TestObj(10, 'abcdef')
        self.assertTrue(m2m.catch_explode(obj))

    def test_no_catch_explode(self):
        obj = entities.TestObj(10, 'abcdef')
        with self.assertRaises(Exception):
            m2m.no_catch_explode(obj)


if __name__ == '__main__':
    unittest.main()
