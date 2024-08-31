# This file runs the Golang PID Simulation in terminal

# For a visual chart of the simulation, visit: https://github.com/theodryden/TDTechnicalSubmissionFrontEnd

# Terminal View
![Screenshot 2024-08-31 at 21 00 54](https://github.com/user-attachments/assets/dea9a8d4-6dcf-4561-9f39-f220d6d27be2)

# To improve
* I would investigate deeper into methods of tuning such as ziegler nichols method. 
* The aperture consistently increments for a minute until it reaches 100%. So it would be good challenge to see how the simulation handles a decrement towards 0% aperture.

# Project Lessons
* Submodule challenges on Git - Always best to use a different development branch for testing even on smaller projects
* Researching existing PID Controller creations on Github such as: https://github.com/felixge/pidctrl
* Learning more about different packages and libraries such as time and reading documentation
* Learning about PID Controllers and the way they keep things at constant like speed, temperature, humidity to name a few.
