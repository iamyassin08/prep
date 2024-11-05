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

export interface Product {
  ID: number;
  ProfileID: string;
  Title: string;
  Description: string;
  ShortDescription: string;
  Price: number;
  Quantity: number;
  DiscountPrice: number;
  RegularPrice: number;
  CreatedAt: string;
  UpdatedAt: string;
  Type: string;
  Thumbnail_url: string;
  Images: Array<{
    ID: number;
    ProductID: number;
    ImageUrl: string;
    CreatedAt: string;
  }>;
  Category: {
    Name: string;
    ID: number;
  };
  ExternalDetails: {
    BrandName: string;
    BrandID: string;
    BrandImageURL: string;
    ExternalURL: string;
  };
}