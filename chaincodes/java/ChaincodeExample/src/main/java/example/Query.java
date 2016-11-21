package example;

import org.hyperledger.java.shim.ChaincodeStub;

/**
 * Created by martin on 10/26/16.
 */
public class Query {

    public String query(ChaincodeStub stub, String function, String[] args) {
        String returnedString = "";

        if(args.length<1){
            return "{\"Error\":\"Incorrect number of arguments. Expecting name of the person to query\"}";
        }

        for(String arg: args) {
            String am =stub.getState(arg);
            if (am!=null&&!am.isEmpty()){
                try{
                    int valA = Integer.parseInt(am);
                    returnedString = returnedString + "{\"Name\":\"" + arg + "\",\"Amount\":\"" + am + "\"}\n";
                }catch(NumberFormatException e ){
                    return "{\"Error\":\"Expecting integer value for asset holding\"}";
                }
            }else{
                return "{\"Error\":\"Failed to get state for " + arg + "\"}";
            }
        }
        return returnedString;
    }
}
