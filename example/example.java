
public class Address {
    @JsonProperty("city")
    private String city;

    @JsonProperty("number")
    private Double number;

    @JsonProperty("country")
    private String country;


    public setCity(value String) {
            this.city = value;
    }
    public getCity() {
            return this.city;
    }

    public setNumber(value Double) {
            this.number = value;
    }
    public getNumber() {
            return this.number;
    }

    public setCountry(value String) {
            this.country = value;
    }
    public getCountry() {
            return this.country;
    }

}
public class PersonalInfo {
    @JsonProperty("hobby")
    private List<String> hobby;

    @JsonProperty("pet_name")
    private String pet_name;


    public setHobby(value List<String>) {
            this.hobby = value;
    }
    public getHobby() {
            return this.hobby;
    }

    public setPet_name(value String) {
            this.pet_name = value;
    }
    public getPet_name() {
            return this.pet_name;
    }

}
public class Person {
    @JsonProperty("name")
    private String name;

    @JsonProperty("personal_info")
    private PersonalInfo personal_info;

    @JsonProperty("nicknames")
    private List<String> nicknames;

    @JsonProperty("addresses")
    private List<Address> addresses;


    public setName(value String) {
            this.name = value;
    }
    public getName() {
            return this.name;
    }

    public setPersonal_info(value PersonalInfo) {
            this.personal_info = value;
    }
    public getPersonal_info() {
            return this.personal_info;
    }

    public setNicknames(value List<String>) {
            this.nicknames = value;
    }
    public getNicknames() {
            return this.nicknames;
    }

    public setAddresses(value List<Address>) {
            this.addresses = value;
    }
    public getAddresses() {
            return this.addresses;
    }

}
