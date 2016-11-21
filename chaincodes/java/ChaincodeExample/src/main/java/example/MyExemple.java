/*
Copyright DTCC 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

         http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package example;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.hyperledger.java.shim.ChaincodeBase;
import org.hyperledger.java.shim.ChaincodeStub;


/**
 * <h1>Classic "transfer" sample chaincode</h1>
 * (java implementation of <A href="https://github.com/hyperledger/fabric/blob/master/examples/chaincode/go/chaincode_example02/chaincode_example02.go">chaincode_example02.go</A>)
 * @author Sergey Pomytkin spomytkin@gmail.com
 *
 */
public class MyExemple extends ChaincodeBase {
	private static Log log = LogFactory.getLog(MyExemple.class);

	@Override
	public String run(ChaincodeStub stub, String function, String[] args) {

		log.info("In run, function:"+function);
        Deploy deploy = new Deploy();
		Invoke invoke = new Invoke();

		switch (function) {
		case "init":
			return deploy.init(stub, function, args);
		case "transfer":
			String re = invoke.transfer(stub, args);
			System.out.println(re);
			return re;
		case "fusion":
			re = invoke.fusion(stub,args);
			System.out.println(re);
			return re;
		case "put":
			for (int i = 0; i < args.length; i += 2)
				stub.putState(args[i], args[i + 1]);
			break;
		case "del":
			for (String arg : args)
				stub.delState(arg);
			break;
		default: 
			return invoke.transfer(stub, args);
		}
	 
		return null;
	}

	@Override
	protected void finalize() throws Throwable {
		super.finalize();
	}

	@Override
	public String query(ChaincodeStub stub, String function, String[] args) {
        Query query = new Query();
        return query.query(stub,function,args);
	}

	@Override
	public String getChaincodeID() {
		return "MyExemple";
	}

	public static void main(String[] args) throws Exception {
		new MyExemple().start(args);
	}
}
