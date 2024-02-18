Shall-Statements:

**S1: A new user shall be able to create an account (FR)** 

Description: A driver provides details such as their name, date of birth, license plate, email and a password to create an account allowing them to access the system and its various functions (e.g. requesting spaces, reviewing their bookings, sending their coordinates to an admin etc.).

**S2: A driver shall be able to send requests for and book parking spaces (FR)**

Description: A driver enters their desired location and time of arrival into the system, which is then reviewed by the admin. Their request is either accepted or rejected. If it’s accepted, the driver is allocated a space, and its details are sent to them. This gives the driver a designated space, so they know exactly where to go and when. If their request is rejected however, the driver knows that they cannot park in the car park at that time.

**S2.1 A driver shall be able to pick their desired location and time, with a drop-down menu that has clear labelling and grouping of options to assist with navigation and selection (NFR)**

Description: When choosing a desired location and time, a driver can use drop down menus which will have the locations of different car parks around campus as well as hours, days, months and years. The clear labelling and grouping of options mean that drivers will not waste time trying to understand the UI, so they can start booking quickly.

**S2.2 This menu shall detect any incorrect dates and times selected (FR)**

Description: The date and time that a driver chooses will be examined by the system to check if it is valid. If it is not, text will appear above the boxes informing them of the error that has occurred (e.g. a driver has chosen a day that has passed). The driver will then have to enter them again. This ensures that all booked spaces have a valid date and time that can be monitored.

**S3 A driver shall be able to view their booked spaces (FR)**

Description: A driver can view all their booked spaces on their profile, with details such as their arrival time. This allows them to ensure there are no issues with their details and so they know exactly when they should arrive to their space.

**S4 A driver shall be able to receive their parking space location along with a map from the main gate (FR)**

Description: Once a driver has arrived at the car park gate, they should receive a map of all parking spaces that tells them where their space is and how to get to it. This ensures that the driver can find their space quickly and does not have to waste time figuring out where it is for themselves.

**S5 A driver shall be able to notify the admin of their arrival at a space using their GPS coordinates (FR)**

Description: When a driver arrives at a space, they can send their current GPS coordinates to the admin. This is necessary for keeping the system metrics accurate and allows the admin to see whether the driver has arrived at the start of their allocated time.

**S6 A driver shall be able to notify the admin when they are leaving campus (FR)**

Description: Before a driver leaves, they can let the admin know that they are doing so. This is necessary for keeping the system metrics accurate and allows the admin to see whether the driver has left within their allocated time.

**S7 A driver shall be clearly informed of the need to notify the admin of their arrival or departure (NFR)**

Description: Drivers will be told of their need to inform the admin when they have arrived or departed their space, and the buttons to do so will be clearly displayed on their space details. A pop-up reminder will appear when the driver uses the system close to the time of their scheduled arrival/departure so they don’t forget. This is to decrease the chances of a driver facing any sort of potential action/follow up from the admin for not informing them that they had arrived/left.

**S8 The admin shall be able to accept or reject parking space requests (FR)**

Description: The admin can look at requests for spaces they have received. The admin then either finds a space themselves or lets the system allocate it. The details of the space are then sent to the driver. This lets the driver know whether their space request has been successful or not and allows for their profile to be updated with their booked space and its details (if their request was successful).

**S8.1 The system shall provide a space to the admin within two seconds of searching for one (NFR)**

Description: When the system searches for a space it should provide a result within two seconds. This allows for faster responses to space requests.

**S9: The admin shall be able to add, remove, block and reserve parking spaces (FR)**

Description: The admin can add new spaces for the booking system, remove spaces from it, block spaces so they cannot be booked and reserve spaces for certain individuals. Certain actions will update the metrics. This ensures that the layout of the car park is accurately reflected in the spaces that can be booked. This prevents conflicting or invalid bookings (e.g. booking a space that has been blocked or reserved) and ensures all available spaces are able to be booked.

**S10: The admin shall be able to monitor the status of the car park with metrics (FR)**

Description: The admin can see a map of the car park which displays taken, available, blocked, reserved, disabled and EV spaces along with metrics of taken and available spaces. This provides the admin with useful information about the car park, allowing them to track its occupancy over time.

**S10.1 The metrics shall be updated in real time (NFR)**

Description: The metrics for the car park (available spaces, unavailable spaces, total rejections etc.) will be updated in real time (around 10ms) to prevent admins or drivers from invalidating each other's actions. This also ensures that the admin has access to accurate information most of the time.

**S11 The drivers and admins shall be able to communicate (FR)**

Description: Drivers and admins can communicate with each other to resolve any issues that drivers might encounter, such as someone else being in their allocated space when they arrive.

**S11.1 Admin and driver communication shall be conducted with a texting interface (NFR)**

Description: Admins and drivers will communicate with each other using a texting interface, as this is a familiar UI for users. This allows for drivers to quickly communicate their problems to the admin without wasting time on getting used to the UI.

**S12 The admin shall be alerted when the GPS coordinates received from a driver and the
allocated parking space do not match (FR)**

Description: The system logs the GPS coordinates of all drivers who have arrived at the car park (as these are uploaded by drivers when they arrive). If it sees a discrepancy between a driver’s coordinates and their current space, the admin is alerted. This ensures that the admin has confirmation that a discrepancy has occurred so they can take any action/ follow-up they deem necessary.

**S13 The admin shall be alerted when a driver’s arrival notification has not been received within one hour from the booked arrival time (FR)**

Description: The admin will be informed when a driver’s GPS coordinates have not been given to them within an hour of the driver’s allocated time. This ensures that the admin has confirmation that the driver has done this so they can take any action/ follow-up they deem necessary.

**S14 An admin shall be alerted when a driver stays one hour or more after the booked departure time (FR)**

Description: The admin will be informed when it has been one hour after a driver’s allocated time has elapsed and they have not received a notification from them to say that they have left. This ensures that the admin has confirmation that the driver has done this so they can take any action/ follow-up they deem necessary.

**S15 The admins and drivers shall have different GUIs (NFR)**

Description: Admins and drivers should have different GUIs which give them access to different features of the system. This is so they can be clearly distinguished and their most important features can be highlighted, such as approving or rejecting requests for admins, and viewing booked spaces for drivers.

**S16 The system shall give the user control over the font size, using an accessibility menu, to help visually impaired users (NFR)**

Description: Users can control the font size, making it bigger or smaller if they need it to be. This will make our system easier to use for visually impaired users, increasing our total user count.
