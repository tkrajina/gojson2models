
public class Address {
    @JsonProperty("city")
    private String city;

    @JsonProperty("number")
    private Double number;

    @JsonProperty("country")
    private String country;


    public void setCity(String value) {
            this.city = value;
    }
    public String getCity() {
            return this.city;
    }

    public void setNumber(Double value) {
            this.number = value;
    }
    public Double getNumber() {
            return this.number;
    }

    public void setCountry(String value) {
            this.country = value;
    }
    public String getCountry() {
            return this.country;
    }

}
public class PersonalInfo {
    @JsonProperty("hobby")
    private ArrayList<String> hobby;

    @JsonProperty("pet_name")
    private String pet_name;


    public void setHobby(ArrayList<String> value) {
            this.hobby = value;
    }
    public ArrayList<String> getHobby() {
            return this.hobby;
    }

    public void setPet_name(String value) {
            this.pet_name = value;
    }
    public String getPet_name() {
            return this.pet_name;
    }

}
public class Person {
    @JsonProperty("name")
    private String name;

    @JsonProperty("personal_info")
    private PersonalInfo personal_info;

    @JsonProperty("nicknames")
    private ArrayList<String> nicknames;

    @JsonProperty("addresses")
    private ArrayList<Address> addresses;


    public void setName(String value) {
            this.name = value;
    }
    public String getName() {
            return this.name;
    }

    public void setPersonal_info(PersonalInfo value) {
            this.personal_info = value;
    }
    public PersonalInfo getPersonal_info() {
            return this.personal_info;
    }

    public void setNicknames(ArrayList<String> value) {
            this.nicknames = value;
    }
    public ArrayList<String> getNicknames() {
            return this.nicknames;
    }

    public void setAddresses(ArrayList<Address> value) {
            this.addresses = value;
    }
    public ArrayList<Address> getAddresses() {
            return this.addresses;
    }

}
