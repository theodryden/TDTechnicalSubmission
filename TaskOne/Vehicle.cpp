#include <iostream>
#include <string>
using namespace std;

/*
--Task--
All classes should have make and wheels,
Motors class inherits make and weels from Vehicle but gets typeOfEngine check and motorStatus condition check (on/off),
Aircraft class inherits all classes and has a unique takeoff() method;
*/

// Base class
class Vehicle {
  public:
    string make;
    int wheels;

    Vehicle(string make, int wheels)
        : make(make), wheels(wheels){};

    int GetWheels(){
        return wheels;
    }  

};
// Derived Wheeled Class 
class Wheeled : public Vehicle {
    public:
        Wheeled(string make, int wheels)
            : Vehicle(make, wheels){}
};
// Motorised Class
class Motorised : public Wheeled {
    public:
        string typeOfEngine;
        bool motorStatus;
        
        Motorised(string make, int wheels, string typeOfEngine, bool motorStatus) 
            // Inherit make and wheels from the vehicle class
            : Wheeled(make, wheels), typeOfEngine(typeOfEngine), motorStatus(motorStatus){
                cout << "Wroom!" << "\n";
            };

        bool switchOn(){
            if(motorStatus){
                cout << "The motor status is currently on!" << "\n";
                return true;
                // There might be the edge case of the motor being on idle
            } else {
                cout << "The motor status is currently off!" << "\n";
                return false;
            }
        };
};
// Aircraft Class
class Aircraft : public Motorised {
    public:
        Aircraft(string make, int wheels, string typeOfEngine, bool motorStatus)
            : Motorised(make, wheels, typeOfEngine, motorStatus){
            };
    
    void takeOff(){
        cout << "The vehicle make is: " << make << "\n";
        cout << "The number of wheels are: " << wheels << "\n";
        cout << "The type of engine is: " << typeOfEngine << "\n";
        cout << "The motor status is: " << motorStatus << "\n";
    }
};
