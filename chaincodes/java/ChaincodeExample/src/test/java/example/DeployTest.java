package example;

import org.hyperledger.java.shim.ChaincodeStub;
import org.junit.Test;

import static junit.framework.Assert.assertEquals;
import static org.junit.Assert.*;
import static org.mockito.Mockito.mock;

/**
 * Created by martin on 10/26/16.
 */
public class DeployTest {

    Deploy deploy = new Deploy();
    ChaincodeStub stubMock = mock(ChaincodeStub.class);

    @Test
    public void whenInit_ShouldTrue() {
        deploy.init(stubMock, "init", new String[]{"a","100","b","200"});
    }

    @Test
    public void whenInit_withTooManyArguments_ShouldError() {
        String response= deploy.init(stubMock,"init",new String[]{"a","100","b","200","c"});
        assertEquals("{\"Error\":\"Incorrect number of arguments. Expecting 4\"}",response);
    }

    @Test
    public void whenInit_withTooFewArguments_ShouldError() {
        String response = deploy.init(stubMock,"init", new String[]{"a","100"});
        assertEquals("{\"Error\":\"Incorrect number of arguments. Expecting 4\"}",response);
    }

    @Test
    public void whenInit_withInvalidArgumentFormat_ShouldError() {
        String response = deploy.init(stubMock,"init",new String[]{"a","c","b","100"});
        assertEquals("{\"Error\":\"Expecting integer value for asset holding\"}",response);
    }

}