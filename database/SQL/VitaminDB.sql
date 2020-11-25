DROP DATABASE IF EXISTS [VitaminDB]
GO
CREATE DATABASE VitaminDB

GO

USE VitaminDB

GO


CREATE TABLE Vitamin
(
	VitaminID INT IDENTITY (1,1),
	VitaminType VARCHAR(32)
	PRIMARY KEY (VitaminID)
)

GO

CREATE TABLE Benefit
(
	BenefitID INT IDENTITY (1,1),
	VitaminID INT,
	Benefit VARCHAR(MAX),
	PRIMARY KEY (BenefitID),
	FOREIGN KEY (VitaminID) REFERENCES Vitamin(VitaminID),

)

GO

CREATE TABLE Inventory (
    [InventoryID] INT IDENTITY (1,1),
    [Name] VARCHAR(50),
    [Count] INT,
    [Site] VARCHAR(50)
	PRIMARY KEY ([InventoryID])
)

GO

CREATE TABLE InventoryVitamin
(
	InventoryVitaminID INT IDENTITY (1,1),
	InventoryID INT,
	VitaminID INT,
	PercentDailyValue INT,
	PRIMARY KEY (InventoryVitaminID),
	FOREIGN KEY (InventoryID) REFERENCES Inventory(InventoryID),
	FOREIGN KEY (VitaminID) REFERENCES Vitamin(VitaminID)
)

