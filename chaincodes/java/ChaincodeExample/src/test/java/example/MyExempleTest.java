package example;

import org.hyperledger.java.shim.ChaincodeStub;
import org.junit.After;
import org.junit.Assert;
import org.junit.Test;
import org.mockito.Mockito;


import java.util.Arrays;

import static junit.framework.Assert.assertEquals;
import static org.mockito.Matchers.any;
import static org.mockito.Matchers.eq;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

/**
 * Created by martin on 10/20/16.
 */
public class MyExempleTest {

    MyExemple myExemple = new MyExemple();

    ChaincodeStub stubMock = mock(ChaincodeStub.class);


    /*@Before
    public void setUp() throws Exception {

    }*/


    @Test
    public void whenInit_ShouldTrue() {
        myExemple.run(stubMock,"init", new String[]{"a","100","b","200"});
    }

    @Test
    public void whenTransfer_ShouldTrue() {
        when(stubMock.getState(eq("a"))).thenReturn("200");
        when(stubMock.getState(eq("b"))).thenReturn("100");
        myExemple.run(stubMock,"transfer",new String[]{"a","b","100"});
    }

//
//    @Test
//    public void queryOK() {
//        stub.putState('a',100);
//        stub.putState('b',200);
//        String str = myExemple.query(stub,'query',['a','b']);
//        assertEquals("{\"Name\":\"a\",\"Amount\":\"100\"}\n{\"Name\":\"b\",\"Amount\":\"200\"}\n",str);
//        stub.delState('a');
//        stub.delState('b');
//    }
//
//    @Test
//    public void fusionOK() {
//        stub.putState('a',100);
//        stub.putState('b',200);
//        myExemple.run(stub,'fusion',['a','b']);
//        assertEquals(300,stub.getState('a'));
//        assertEquals(null,stub.getState('b'));
//        stub.delState('a');
//    }

}