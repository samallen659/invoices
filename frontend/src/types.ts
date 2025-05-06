type InvoiceStatus = "pending" | "paid" | "draft";

type Address = {
  Street: string;
  City: string;
  PostCode: string;
  Country: string;
};

type Item = {
  Name: string;
  Quantity: number;
  Price: number;
  Total: number;
};

type ItemRes = {
  Name: string;
  Price: number;
};

type InvoiceItem = {
  Item: ItemRes;
  Quantity: number;
  Total: number;
};

type Client = {
  ClientName: string;
  ClientEmail: string;
};

type Invoice = {
  ID: string;
  CreatedAt: string;
  PaymentDue: string;
  Description: string;
  PaymentTerms: number;
  ClientName: string;
  ClientEmail: string;
  Status: InvoiceStatus;
  ClientAddress: Address;
  SenderAddress: Address;
  Items: Item[];
  Total: number;
};

type InvoiceRes = {
  ID: string;
  CreatedAt: Date;
  PaymentDue: Date;
  Description: string;
  PaymentTerms: number;
  Client: Client;
  Status: InvoiceStatus;
  ClientAddress: Address;
  SenderAddress: Address;
  InvoiceItems: InvoiceItem[];
  Total: number;
};

type User = {
  UserName: string;
  FirstName: string;
  LastName: string;
  Email: string;
};

type LoginRequest = {
  UserName: string;
  Email: string;
  Password: string;
};

type SignUpRequest = {
  UserName: string;
  FirstName: string;
  LastName: string;
  Email: string;
  Password: string;
};

type GetUserRequest = {
  UserName: string;
  Email: string;
};

export type {
  InvoiceStatus,
  Address,
  Item,
  InvoiceItem,
  Client,
  Invoice,
  InvoiceRes,
  ItemRes,
  User,
  LoginRequest,
  SignUpRequest,
  GetUserRequest,
};
