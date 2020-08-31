DROP DATABASE IF EXISTS [VitaminDB]
GO
CREATE DATABASE VitaminDB

GO

USE VitaminDB

GO


CREATE TABLE Vitamin
(
	VitaminID INT IDENTITY (1,1),
	VitaminType VARCHAR(32),
	Benefits VARCHAR(MAX)
	PRIMARY KEY (VitaminID)
)

CREATE TABLE HeartHealth
(
	HeartHealthID INT IDENTITY (1,1) ,
	VitaminID INT NOT NULL,
	CONSTRAINT PK_HeartHealthID PRIMARY KEY (HeartHealthID),
	CONSTRAINT FK_HeartHealth_Vitamin_VitaminID FOREIGN KEY (VitaminID) REFERENCES Vitamin(VitaminID)
)

GO

