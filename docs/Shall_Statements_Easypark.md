Shall-Statements:

**S1: A new user shall be able to create an account (FR).**

Description: The driver provides details such as their name, phone number, email, approximate location and (optionally) their card details to create an account allowing the driver to access the system.

**S2: A driver shall be able to send requests for and book parking spaces (FR)**

Description: The driver enters their desired location and time of arrival into Easypark, which is then reviewed by the admin. It is either accepted or rejected. If it’s accepted, the driver is given the details of the space. This is a core feature of the system.

**S2.1 A driver shall be able to pick their desired location and time with a drop-down menu (NFR)**

Description: When choosing a desired location and time, the driver can use drop down menus which will have the locations of different car parks around campus as well as hours, days, months and years. This is a simple UI to use.

**S2.2 This menu shall detect any anomalies in dates and times selected (FR)**

Description: The date and time that the user chooses will be examined by the system to check if it is valid. If it is not, the system will reject the details and make the user enter them again. This is to prevent invalid dates from being recorded in the system.

**S3 A driver shall be able to view their booked spaces (FR)**

Description: A driver can view all their booked spaces on their profile, with details such as their arrival time. This allows them to keep track of them to ensure there are no issues with their details.

**S4 A driver shall be able to receive their parking space location along with a map from the main gate (FR)**

Description: Once a driver has arrived at the car park gate, they should receive a map of all parking spaces that tells them where their space is and how to get to it. This improves user experience for the user.

**S5 A driver shall be able to notify admins of their arrival at a space using their GPS coordinates (FR)**

Description: When a driver arrives at a space, they can send their current GPS coordinates to admins so they can see whether they are in the right location and so the system metrics can be updated. This also allows for driver tracking to ensure that they adhere to their space’s allocated time.

**S6 A driver shall be able to notify admins when they are leaving campus (FR)**

Description: Before a driver leaves, they can let admins know that they are doing so. This ensures that drivers stay for only their allocated time and allows system metrics to be updated. And keeps the admins up to date with the current state of the car park.

**S7 A driver shall be clearly informed of the need to notify the admin of their arrival or departure (NFR)**

Description: The driver will be told of their need to inform the admin when they have arrived or departed their space, and the buttons to do so will be clearly displayed on their space details.

**S8 An admin shall be able to accept and reject parking space requests (FR)**

Description: The admin can look at the driver requests for spaces they have received. The admin then either finds a space themselves or lets the system allocate it. The details of the space are then sent to the user. Finding and sending space details is a core functionality of the system.

**S8.1 The system shall provide a space to the admin within two seconds of searching for one (NFR)**

Description: When the system searches for a space it should provide a result within two seconds. This allows for faster system operations.

**S9: An admin shall be able to add, remove, block and reserve parking spaces (FR)**

Description: An admin can add new spaces for the booking system, remove spaces from it, block spaces so they cannot be booked and reserve spaces for certain individuals. Certain actions will update the metrics. These are a few of the core features of the system.

**S10: An admin shall be able to monitor the status of the car park with metrics (FR)**

Description: An admin can see a map of their car park which displays taken, available, blocked, reserved, disabled and EV spaces along with metrics of taken and available spaces. Provides the admin with useful information about the parking lots.

**S10.1 The metrics shall be updated in real time (NFR)**

Description: The metrics for the car park (available spaces, unavailable spaces, total rejections etc.) will be updated in real time (around 10ms) to prevent admins or drivers from invalidating each other's actions.

**S11 The drivers and admins shall be able to communicate (FR)**

Description: Drivers and admins can communicate with each other to resolve any issues that the driver might encounter.

**S11.1 This admin and driver communication shall be conducted with a texting interface (NFR)**

Description: Admins and drivers will communicate with each other using a texting interface, as this is a familiar UI for users.

**S12 The admin shall be alerted when the GPS coordinates received from a driver and the allocated parking space do not match (FR)**

Description: The system logs the GPS coordinates of all drivers who have arrived at the car park (as these are uploaded by drivers when they arrive). If it sees a discrepancy between the driver’s coordinates and their current space, the admin is alerted.

**S13 The admin shall be alerted when a driver’s arrival notification has not been received within one hour from the booked arrival time (FR)**

Description: The admin will be informed when a user’s GPS coordinates have not been given to them within an hour of that driver’s allocated time. This allows the admin to take any action/ follow-up if they deem necessary and to keep metrics up to date.

**S14 The admin shall be alerted when a driver stays one hour or more after the booked departure time (FR)**

Description: The admin will be informed when it has been one hour after a driver’s allocated time has elapsed and they have not received a notification from them to say that they have left. This is so the system knows that a driver has left, after which the metrics can be updated

**S15 The admins and drivers should have different GUIs (NFR)**

Description: Admins and drivers should have different GUIs which give them access to different features of the system. 

**S16 The system shall have various accessibility features such as magnification, changing font size and high contrast (NFR)**

Description: Users will be able to use a range of accessibility features such as magnification, adjustable font size and high contrast. This will enable our website to be accessible more users.
