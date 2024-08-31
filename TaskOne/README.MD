# Lessons
* Inheritance in c++
- Constructors in c++
- Building files in c++
- Compiling files in c++
- Creating header files
- The file you’re on matters when entering the build command
- to_string method to convert the wheel int to string
- **Inheritance Access Specifiers**:
    - Public keyword to ensure that the members of `Vehicle` remain accessible to `Motorised` and `Aircraft`.
- **Constructor Initialization**:
    - So you need the constructor on all classes
    - In C++, when a class inherits from multiple classes that each inherit from a common base class virtually, the derived class must explicitly initialize the virtual base class.
    - The derived class's constructor must explicitly call the constructor of the virtual base class to set up the base class part of the object. This is necessary because the compiler needs to know exactly how to initialize the virtual base class.
- **Virtual inheritance** - from the “Definition is ambigous error”

# Self-Management
- Write thought process out on paper and pen first
- Identify blockers when brute force isn’t working
- Find examples, practice them from memory working, apply logic to your novel problem
- Take short breaks
- Debug and write the classes/functions in stages
