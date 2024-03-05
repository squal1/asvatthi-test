def add(x, y):
    return x + y


class Car:
    def __init__(self, make, model, year, color):
        self.make = make
        self.model = model
        self.year = year
        self.color = color
        self.is_running = False

    def start(self):
        if not self.is_running:
            self.is_running = True
            print(f"The {self.color} {self.make} {self.model} starts.")
        else:
            print("The car is already running.")

    def stop(self):
        if self.is_running:
            self.is_running = False
            print(f"The {self.color} {self.make} {self.model} stops.")
        else:
            print("The car is already stopped.")
