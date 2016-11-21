package example;

import org.hyperledger.java.shim.ChaincodeStub;

/**
 * Created by martin on 10/26/16.
 */
public class Invoke {

    public String transfer(ChaincodeStub stub, String[] args) {
        System.out.println("in transfer");
        if(args.length!=3){
            System.out.println("Incorrect number of arguments:"+args.length);
            return "{\"Error\":\"Incorrect number of arguments. Expecting 3: from, to, amount\"}";
        }
        String fromName =args[0];
        String fromAm=stub.getState(fromName);
        String toName =args[1];
        String toAm=stub.getState(toName);
        String am =args[2];
        int valFrom=0;
        if (fromAm!=null&&!fromAm.isEmpty()){
            try{
                valFrom = Integer.parseInt(fromAm);
            }catch(NumberFormatException e ){
                System.out.println("{\"Error\":\"Expecting integer value for asset holding of "+fromName+" \"}"+e);
                return "{\"Error\":\"Expecting integer value for asset holding of "+fromName+" \"}";
            }
        }else{
            return "{\"Error\":\"Failed to get state for " +fromName + "\"}";
        }

        int valTo=0;
        if (toAm!=null&&!toAm.isEmpty()){
            try{
                valTo = Integer.parseInt(toAm);
            }catch(NumberFormatException e ){
                e.printStackTrace();
                return "{\"Error\":\"Expecting integer value for asset holding of "+toName+" \"}";
            }
        }else{
            return "{\"Error\":\"Failed to get state for " +toName + "\"}";
        }

        int valA =0;
        try{
            valA = Integer.parseInt(am);
        }catch(NumberFormatException e ){
            e.printStackTrace();
            return "{\"Error\":\"Expecting integer value for amount \"}";
        }
        if(valA>valFrom)
            return "{\"Error\":\"Insufficient asset holding value for requested transfer amount \"}";
        valFrom = valFrom-valA;
        valTo = valTo+valA;
        System.out.println("Transfer "+fromName+">"+toName+" am='"+am+"' new values='"+valFrom+"','"+ valTo+"'");
        stub.putState(fromName,""+ valFrom);
        stub.putState(toName, ""+valTo);

        System.out.println("Transfer complete");

        return null;

    }

    public String fusion(ChaincodeStub stub, String[] args){
        if(args.length<2){
            return "{\"Error\":\"Incorrect number of arguments. Expecting name of the person to query\"}";
        }

        String returnedString = "Vous avez fusionné les comptes";
        int total = 0;
        for(String arg : args){
            String am = stub.getState(arg);
            if(am != null && !am.isEmpty()){
                total = total + Integer.parseInt(am);
                returnedString = returnedString + " " + arg;
            }else {
                System.out.println("{\"Error\":\"Failed to get state for " + arg + "\"}");
                return "{\"Error\":\"Failed to get state for " + arg + "\"}";
            }
        }
        for(String arg : args){
            stub.delState(arg);
        }
        stub.putState(args[0], "" + total);
        returnedString = returnedString + " dans le compte " + args[0];
        System.out.println(returnedString);
        return returnedString;
    }
}
