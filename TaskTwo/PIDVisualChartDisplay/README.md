# These files contain the Golang code to be run injunction with the React Front End
* Clone this repository
* Run this Golang Simulation
* The data will be sent from the Golang simulator to the React website display
* (The React listens for data being sent to port: 8080 and displays the data on the standard localhost: 3000)

# Chart Golang simulation: https://github.com/theodryden/TDTechnicalSubmissionFrontEnd

# To improve
* I would investigate deeper into methods of tuning such as ziegler nichols method. 
* The aperture consistently increments for a minute until it reaches 100%. So it would be good challenge to see how the simulation handles a decrement towards 0% aperture.

# Project Lessons
* Submodule challenges on Git - Always best to use a different development branch for testing even on smaller projects
* Researching existing PID Controller creations on Github such as: https://github.com/felixge/pidctrl
* Learning more about different packages and libraries such as time and reading documentation
* Learning about PID Controllers and the way they keep things at constant like speed, temperature, humidity to name a few.
