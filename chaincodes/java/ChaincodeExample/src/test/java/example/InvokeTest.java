package example;

import org.hyperledger.java.shim.ChaincodeStub;
import org.junit.Test;

import static junit.framework.Assert.assertEquals;
import static org.junit.Assert.*;
import static org.mockito.Matchers.eq;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

/**
 * Created by martin on 10/26/16.
 */
public class InvokeTest {

    Invoke invoke = new Invoke();
    ChaincodeStub stubMock = mock(ChaincodeStub.class);

    @Test
    public void whenTransfer_ShouldTrue() {
        when(stubMock.getState(eq("a"))).thenReturn("200");
        when(stubMock.getState(eq("b"))).thenReturn("100");
        invoke.transfer(stubMock, new String[]{"a","b","100"});
    }

    @Test
    public void whenTransfer_withoutSufficientAssetHolding_ShouldError() {
        when(stubMock.getState(eq("a"))).thenReturn("1");
        when(stubMock.getState(eq("b"))).thenReturn("2");
        String response = invoke.transfer(stubMock, new String[]{"a","b","100"});
        assertEquals(response, "{\"Error\":\"Insufficient asset holding value for requested transfer amount \"}");
    }

    @Test
    public void whenTransfer_withTooManyArguments_ShouldError() {
        when(stubMock.getState(eq("a"))).thenReturn("200");
        when(stubMock.getState(eq("b"))).thenReturn("100");
        String response= invoke.transfer(stubMock, new String[]{"a","b","100","c"});
        assertEquals("{\"Error\":\"Incorrect number of arguments. Expecting 3: from, to, amount\"}",response);
    }

    @Test
    public void whenTransfer_withTooFewArguments_ShouldError() {
        when(stubMock.getState(eq("a"))).thenReturn("200");
        String response= invoke.transfer(stubMock, new String[]{"a","100"});
        assertEquals("{\"Error\":\"Incorrect number of arguments. Expecting 3: from, to, amount\"}",response);
    }

}