
class Address {
    city: string;
    number: number;
    country: string;
}
class PersonalInfo {
    hobby: string[];
    pet_name: string;
}
class Person {
    name: string;
    personal_info: PersonalInfo;
    nicknames: string[];
    addresses: Address[];
}
