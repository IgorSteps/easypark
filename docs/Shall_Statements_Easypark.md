Shall-Statements:

S1: A new user shall be able to create an account (FR) 

Description: The driver provides details such as their name, phone number, email, approximate location and (optionally) their card details to create an account.

S2: A driver shall be able to send requests for and book parking spaces (FR)

Description: The driver enters their desired location and time of arrival into Easypark, which it uses to find a space and present it to the driver for booking and payment (if one can be found). This booking is then either rejected or approved by the admin

S2.1 A driver shall be able to pick their desired location and time with a drop-down menu (NFR)

Description: When choosing a desired location and time, the user can use drop down menus which will have the locations of different car parks around campus as well as hours, days, months and years

S3 A driver shall be able to view their booked spaces (FR)

Description: A driver can view all their booked spaces on their profile, with details such as their arrival time

S4 A driver shall be able to receive their parking space location along with a map from the main gate (FR)

Description: Once a driver has arrived at the car park gate, they should receive a map of all parking spaces that tells them where their space is and how to get to it

S5 A driver shall be able to notify admins of their arrival at a space using their GPS coordinates (FR)

Description: When a driver arrives at a space, they can send their current GPS coordinates to admins so they can see whether they are in the right location and so the system metrics can be updated

S6 A driver shall be able to notify admins when they are leaving campus (FR)

Description: Before a driver leaves, they can let admins know that they are doing so. This ensures that drivers stay for only their allocated time and allows system metrics to be updated

S7: An admin shall be able to accept and reject parking space requests (FR)

Description: An admin reviews requests from drivers for parking spaces and accepts or rejects them, depending on whether the space can be allocated

S8: An admin shall be able to add, remove, block and reserve parking spaces (FR)

Description: An admin can add new spaces for the booking system, remove spaces from it, block spaces so they cannot be booked and reserve spaces for certain individuals

S9: An admin shall be able to monitor the status of the car park with metrics (FR)

Description: An admin can see a map of their car park which displays taken, available, blocked, disabled and EV spaces along with metrics of taken and available spaces

S9.1 The map shall be colour coded with a drag and drop system (NFR)

Description: Different colours will correspond to different space types, with a drag and drop system that can be used to book spaces and block them

S9.2 The metrics shall be updated in real time (NFR)

Description: The metrics for the car park (available spaces, unavailable spaces, total rejections etc.) will be updated in real time

S10 The drivers and admins shall be able to communicate (FR)

Description: Drivers and admins can communicate with each other over a texting interface to resolve any issues the driver has

S11 The admin shall be alerted when the GPS coordinates received from a driver and the
allocated parking space do not match (FR)

Description: The system logs the GPS coordinates of all drivers who have arrived at the car park (as these are uploaded by drivers when they arrive). If it sees a discrepancy between the driver’s coordinates and their current space, the admin is alerted

S12 The admin shall be alerted when a driver’s arrival notification has not been received within one hour from the booked arrival time (FR)

Description: The admin will be informed when a user’s GPS coordinates have not been given to them within an hour of that driver’s allocated time

S13 The admin shall be alerted when a driver stays one hour or more after the booked departure time (FR)

Description: The admin will be informed when it has been one hour after a driver’s allocated time has elapsed and they have not received a notification from them to say that they have left

S14 The admins and drivers should have different GUIs (NFR)

Description: Admins and drivers should have different GUIs which give them access to different features
