package example;

import org.hyperledger.java.shim.ChaincodeStub;

/**
 * Created by martin on 10/26/16.
 */
public class Deploy {

    public String init(ChaincodeStub stub, String function, String[] args) {
        if(args.length!=4){
            return "{\"Error\":\"Incorrect number of arguments. Expecting 4\"}";
        }
        try{
            int valA = Integer.parseInt(args[1]);
            int valB = Integer.parseInt(args[3]);
            stub.putState(args[0], args[1]);
            stub.putState(args[2], args[3]);
        }catch(NumberFormatException e ){
            return "{\"Error\":\"Expecting integer value for asset holding\"}";
        }
        return null;
    }
}
