import re


def main():
  N = int(input())

  user_names = [input() for i in range(N)]
  user_ids = [int(re.sub(r"\D", "", user_names[i])) for i in range(len(user_names))]

  new_dict = sorted(dict(zip(user_ids, user_names)).items())
  
  for i in range(len(new_dict)):
    print(new_dict[i][1])


if __name__ == "__main__":
  main()