class UserLogin {
  password: string = "";
  email: string = "";
}

class UserSignup {
  email: string = "";
  password: string = "";
  passwordConfrim: string = "";
  name: string = "";
}

class User {
  id: string = "";
  email: string = "";
  name: string = "";
  password: string = "";
}

export interface User {
  ID: number;
  UserID: string;
  FirstName: string;
  Description: string;
  ShortDescription: string;
  Email : number;
  Quantity: number;
  DiscountEmail : number;
  LastName : number;
  CreatedAt: string;
  UpdatedAt: string;
  Type: string;
  Thumbnail_url: string;
  Images: Array<{
    ID: number;
    UserID: number;
    ImageUrl: string;
    CreatedAt: string;
  }>;
  Category: {
    FirstName: string;
    ID: number;
  };
  ExternalDetails: {
    BrandFirstName: string;
    BrandID: string;
    BrandImageURL: string;
    ExternalURL: string;
  };
}