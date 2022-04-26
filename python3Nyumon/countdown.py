def main():
    english = "Monday", "Tuesday", "Wednesday"
    french = "lundi", "mardi", "mercredi"

    print(list(zip(english, french)))
    print(dict(zip(english, french)))

if __name__ == "__main__":
    main()
