
public class Models {

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static final class  {
        @JsonProperty("aaa")
        private String aaa;


        public void setAaa(String value) {
                this.aaa = value;
        }
        public String getAaa() {
                return this.aaa;
        }

    }
    @JsonIgnoreProperties(ignoreUnknown = true)
    public static final class Person {
        @JsonProperty("bu")
        private XXX bu;


        public void setBu(XXX value) {
                this.bu = value;
        }
        public XXX getBu() {
                return this.bu;
        }

    }
}
