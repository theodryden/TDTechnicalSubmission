#include <iostream>
#include "Vehicle.h"
#include <string>

using namespace std;




int main() {
  // Vehicle OldBanger("Ford", 4);
  // Aircraft fThirtyFive("Lockheed Martin", 3,"V6",false);

  // Ford object 
  Motorised Mondeo("Ford", 4, "Diesel",false);
  cout << "Your Ford has " + to_string(Mondeo.GetWheels()) + " wheels" << "\n";
  cout << "\n";
  
  // Airline object 
  Aircraft Airliner("Boeing", 3, "Kerosene",true);
  Airliner.switchOn();
  Airliner.takeOff();
  cout << Airliner.make + " " + to_string(Airliner.wheels) + " wheeler" + " " + Airliner.typeOfEngine + " " + "status: " + to_string(Airliner.motorStatus)<< "\n";

}