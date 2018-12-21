import unittest
import random
from string import ascii_letters

str_size = 10
list_size = 50

def rand_string():
  return ''.join(random.sample(ascii_letters, str_size))


def rand_list(size=list_size):
  l = []
  for x in range(size):
    l.append(rand_string())
  return l


class TestDicts(unittest.TestCase):
  def test_dict_order_is_predictable_with_predictable_inputs(self):
    # how many times will we test?
    times = 10
    # get a random list of strings
    r_list = rand_list()
    # ensure the items in our list are in random order
    random.shuffle(r_list)

    # add all of the items from the random list into a dict; add that dict to all dicts
    all_dicts = []
    for _ in range(times):
      d = {}
      for item in r_list:
        d[item] = 'val'
      all_dicts.append(d)

    # ensure we have the number of dicts we expect
    self.assertTrue(len(all_dicts) == times)

    # from a dict for each loop, write all of the keys into a list
    all_lists = []
    for d in all_dicts:
      ordered_key_list = []
      for key in d.keys():
        ordered_key_list.append(key)
      all_lists.append(ordered_key_list)

    # verify all of the dicts are the same
    for l in all_lists:
      self.assertTrue(l == all_lists[0])  # <--- is True, deterministic

