import math


def main():
  input_info = input().split()
  x_list = [int(x) for x in input_info]

  winning_result = []
  for i in range(x_list[0]):
    winning_result.append("N")

  lcm = (x_list[1]*x_list[2])//math.gcd(x_list[1], x_list[2])

  for i in range(len(winning_result)):
    if (i+1) % x_list[1] == 0:
      winning_result[i] = "A"
    if (i+1) % x_list[2] == 0:
      winning_result[i] = "B"
    if (i+1) % lcm == 0:
      winning_result[i] = "AB"
    
  for v in winning_result:
    print(v)

if __name__ == "__main__":
  main()