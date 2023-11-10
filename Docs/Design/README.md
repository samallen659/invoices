<p><a target="_blank" href="https://app.eraser.io/workspace/pk4x6Qf2yW4iIWA0qDis" id="edit-in-eraser-github-link"><img alt="Edit in Eraser" src="https://firebasestorage.googleapis.com/v0/b/second-petal-295822.appspot.com/o/images%2Fgithub%2FOpen%20in%20Eraser.svg?alt=media&amp;token=968381c8-a7e7-472a-8ed6-4a6626da5501"></a></p>



# Project: Invoice App


## Overview 
A web based app that allows the creation, editing and deletion of invoices.

## Domain Model


### Entities


- Invoice
    - ID
    - Created At
    - Payment Due
    - Description
    - PaymentTerms
    - Status
    - Total


- User
    - ID
    - First Name
    - Last Name
    - Email


### Value Objects


- Address
    - Street
    - City
    - Post Code
    - Country


- Item
    - Name
    - Price


- Invoice Item
    - ID
    - Quantity
    - Total


- Client
    - ID
    - Client Name
    - Client Email


### Aggregates


- Invoice (Entity is aggregate root)
    - Client
    - List of Invoice Items
    - Sender Address
    - Client Address 




## Data Model


![Data Model](/.eraser/pk4x6Qf2yW4iIWA0qDis___CPD07GfrQMaqPBhpxFXyqNLI4YC3___---figure---FFksyMooWwd9_4Qiv3Mox---figure---HeBRd7nOO-KInGVeC58OOQ.png "Data Model")




<!--- Eraser file: https://app.eraser.io/workspace/pk4x6Qf2yW4iIWA0qDis --->