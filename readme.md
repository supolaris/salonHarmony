Overview:
This document outlines the requirements for the backend development of a salon
owner interface. The frontend has been developed with functionalities including user
registration, services management, staff management, profile updates, and
appointment handling.
Functionality Overview:

1.  User Registration:
    ● Uponclicking the "Get Started" button, the owner is directed to a
    registration form.
    ● Uponfilling the form, the owner's data needs to be saved in the
    database.
2.  Services Management:
    ● Allows CRUD (Create, Read, Update, Delete) operations on salon
    services.
3.  Staff Management:
    ● Similar to services management, CRUD operations are enabled for
    managing staff details.
4.  Profile Updates:
    ● Ownerscan update their information through a profile page, with
    changes saved to the database.
5.  Appointment Handling:
    ● Theappointment page has three sub-pages: upcoming appointments,
    confirmed appointments, and completed appointments.
    ● Initially, bookings are manually inputted; integration with a booking
    system will be added later.
    Appointment Page Details:
6.  Upcoming Appointments:
    ● Ownerscan view pending appointments here.
    ● Confirmation or rejection of appointments is required.
    ● Confirmed appointments are moved to the "Confirmed Appointments"
    page.
7.  Confirmed Appointments:
    ● Ownerscan mark appointments as completed here.
    ● Uponcompletion, appointments are moved to the "Completed
    Appointments" page.
8.  Completed Appointments:
    ● Ownerscan view past appointments here.
    ● Adelete option is available to remove completed appointments, serving
    as a history log
