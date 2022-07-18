# neve-spring

## annotations
 <table style="width:1344pt">
 <colgroup>
  <col width="166" style="mso-width-source:userset;mso-width-alt:5312;width:125pt"> 
  <col width="198" style="mso-width-source:userset;mso-width-alt:6336;width:149pt"> 
  <col width="201" style="mso-width-source:userset;mso-width-alt:6432;width:151pt"> 
  <col width="132" style="mso-width-source:userset;mso-width-alt:4224;width:99pt"> 
  <col width="118" style="mso-width-source:userset;mso-width-alt:3776;width:89pt"> 
  <col width="138" style="mso-width-source:userset;mso-width-alt:4416;width:104pt"> 
  <col width="72" style="width:54pt"> 
  <col width="249" style="mso-width-source:userset;mso-width-alt:7968;width:187pt"> 
  <col width="514" style="mso-width-source:userset;mso-width-alt:16448;width:386pt"> 
 </colgroup>
 <tbody>
  <tr height="24"> 
   <td class="xl65">category</td> 
   <td class="xl65">annotation</td> 
   <td class="xl65">annotation description</td> 
   <td class="xl65">target</td> 
   <td class="xl65">parameter</td> 
   <td class="xl65">parameter type</td> 
   <td class="xl65">required</td> 
   <td class="xl65">parameter description</td> 
   <td class="xl65">example</td> 
  </tr> 
  <tr height="132"> 
   <td rowspan="7" class="xl64">Dependency injection</td> 
   <td class="xl64">component</td> 
   <td class="xl64">component is used to denote a type as component. It means that neve will autodetect these types for dependency injection.</td> 
   <td class="xl64">1.type(struct)</td> 
   <td class="xl64">value</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">bean name</td> 
   <td class="xl64">// +neve:controllor:value="userhandler"</td> 
  </tr> 
  <tr height="110"> 
   <td class="xl64">service</td> 
   <td class="xl64">serive is used to denote a type privides some services. neve will autodetect these types for dependency injection.</td> 
   <td class="xl64">1.type(struct)</td> 
   <td class="xl64">value</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">bean name</td> 
   <td class="xl64">// +neve:service:value="userservice"</td> 
  </tr> 
  <tr height="110"> 
   <td class="xl64">controllor</td> 
   <td class="xl64">controller used with web applications or REST web services. neve will autodetect these types for dependency injection.</td> 
   <td class="xl64">1.type(struct)</td> 
   <td class="xl64">value</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">bean name</td> 
   <td class="xl64">// +neve:controller:value="usercontroller"</td> 
  </tr> 
  <tr height="51"> 
   <td rowspan="3" class="xl64">bean</td> 
   <td rowspan="3" class="xl64">bean is applied on a function or a method to specify that it returns a bean to be managedd by neve.</td> 
   <td rowspan="3" class="xl64">1.function 2.method</td> 
   <td class="xl64">value</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">bean name</td> 
   <td class="xl64">// +neve:bean:value="usercontroller"</td> 
  </tr> 
  <tr height="59"> 
   <td class="xl64">initmethod</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">specify bean init method (method must be public and have no parameters)</td> 
   <td class="xl64">// +neve:bean:initmethod="Init"</td> 
  </tr> 
  <tr height="53"> 
   <td class="xl64">destroymethod</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">specify bean destroy method (method must be public and have no parameters)</td> 
   <td class="xl64">// +neve:bean:destroymethod="Close"</td> 
  </tr> 
  <tr height="154"> 
   <td class="xl64">scope</td> 
   <td class="xl64">declaring a scope for defining bean</td> 
   <td class="xl64">1.type 2.function 3.method</td> 
   <td class="xl64">value</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">scopes: 1.singleton(default) return the same bean instance each time one is needed. 2.prototype produce a new bean instance each time one is needed.</td> 
   <td class="xl64">// +neve:scope:value="prototype"</td> 
  </tr> 
  <tr height="22"> 
   <td rowspan="22" class="xl64">REST</td> 
   <td rowspan="2" class="xl64">requestmapping</td> 
   <td rowspan="2" class="xl64">used to map web requests onto specific controller types hand methods.</td> 
   <td rowspan="2" class="xl64">1.type 2.method</td> 
   <td class="xl64">value</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">URI pattern</td> 
   <td class="xl64">// +neve:requestmapping:value="/user/:id"</td> 
  </tr> 
  <tr height="98"> 
   <td class="xl64">method</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">http method (must be all letter uppercase)</td> 
   <td class="xl64">// +neve:requestmapping:method="POST"</td> 
  </tr> 
  <tr height="24"> 
   <td rowspan="3" class="xl64">requestparam</td> 
   <td rowspan="3" class="xl64">retrieve the URL query parameters and map it to the method argument.</td> 
   <td rowspan="3" class="xl64">1.method</td> 
   <td class="xl64">name</td> 
   <td class="xl64">string</td> 
   <td class="xl64">TRUE</td> 
   <td class="xl64">method argument name</td> 
   <td rowspan="3" class="xl64">// +neve:requestparam:name="projectId",default="-1",required=false</td> 
  </tr> 
  <tr height="29"> 
   <td class="xl64">default</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">default value</td> 
  </tr> 
  <tr height="92"> 
   <td class="xl64">required</td> 
   <td class="xl64">bool</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">if parameter is missing and without any default value will response with HTTP 400 error(default false)</td> 
  </tr> 
  <tr height="22"> 
   <td rowspan="3" class="xl64">pathvariable</td> 
   <td rowspan="3" class="xl64">map the URI variable to one of the method arguments.</td> 
   <td rowspan="3" class="xl64">1.method</td> 
   <td class="xl64">name</td> 
   <td class="xl64">string</td> 
   <td class="xl64">TRUE</td> 
   <td class="xl64">method argument name</td> 
   <td rowspan="3" class="xl64">// +neve:pathvariable:name="userId",default="-1",required=true</td> 
  </tr> 
  <tr height="22"> 
   <td class="xl64">default</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">default value</td> 
  </tr> 
  <tr height="88"> 
   <td class="xl64">required</td> 
   <td class="xl64">bool</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">if parameter is missing and without any default value will response with HTTP 400 error(default false)</td> 
  </tr> 
  <tr height="22"> 
   <td rowspan="3" class="xl64">requestheader</td> 
   <td rowspan="3" class="xl64">map the request Header variable to one of the method arguments.</td> 
   <td rowspan="3" class="xl64">1.method</td> 
   <td class="xl64">name</td> 
   <td class="xl64">string</td> 
   <td class="xl64">TRUE</td> 
   <td class="xl64">method argument name</td> 
   <td rowspan="3" class="xl64">// +neve:requestheader:name="clientid",default="-1"</td> 
  </tr> 
  <tr height="22"> 
   <td class="xl64">default</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">default value</td> 
  </tr> 
  <tr height="88"> 
   <td class="xl64">required</td> 
   <td class="xl64">bool</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">if parameter is missing and without any default value will response with HTTP 400 error(default false)</td> 
  </tr> 
  <tr height="22"> 
   <td rowspan="2" class="xl64">requestbody</td> 
   <td rowspan="2" class="xl64">map the request body to one of the method arguments.</td> 
   <td rowspan="2" class="xl64">1.method</td> 
   <td class="xl64">name</td> 
   <td class="xl64">string</td> 
   <td class="xl64">TRUE</td> 
   <td class="xl64">method argument name</td> 
   <td rowspan="2" class="xl64">// +neve:requestbody:name="clientid"</td> 
  </tr> 
  <tr height="88"> 
   <td class="xl64">required</td> 
   <td class="xl64">bool</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">if parameter is missing and without any default value will response with HTTP 400 error(default false)</td> 
  </tr> 
  <tr height="44"> 
   <td rowspan="5" class="xl64">loghttp</td> 
   <td rowspan="5" class="xl64">output http logs</td> 
   <td rowspan="5" class="xl64">1.method</td> 
   <td class="xl64">norequestheader</td> 
   <td class="xl64">bool</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">output http logs without request header</td> 
   <td rowspan="5" class="xl64">// +neve:loghttp:noresponsebody=true</td> 
  </tr> 
  <tr height="44"> 
   <td class="xl64">norequestbody</td> 
   <td class="xl64">bool</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">output http logs without request body</td> 
  </tr> 
  <tr height="44"> 
   <td class="xl64">noresponseheader</td> 
   <td class="xl64">bool</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">output http logs without response header</td> 
  </tr> 
  <tr height="44"> 
   <td class="xl64">noresponsebody</td> 
   <td class="xl64">bool</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">output http logs without response body</td> 
  </tr> 
  <tr height="44"> 
   <td class="xl64">level</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">log leve (options:debug | info | warn | error) default into</td> 
  </tr> 
  <tr height="44"> 
   <td rowspan="4" class="xl64">swagger:apioperation</td> 
   <td rowspan="4" class="xl64">enable swagger</td> 
   <td rowspan="4" class="xl64">1.method</td> 
   <td class="xl64">value</td> 
   <td class="xl64">string</td> 
   <td class="xl64">TRUE</td> 
   <td class="xl64">A short summary of what the operation does</td> 
   <td rowspan="4" class="xl64">// +neve:swagger:apioperation:value="create user"</td> 
  </tr> 
  <tr height="44"> 
   <td class="xl64">notes</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">A verbose explanation of the operation behavior</td> 
  </tr> 
  <tr height="66"> 
   <td class="xl64">tags</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">A list of tags to each API operation that separated by commas</td> 
  </tr> 
  <tr height="44"> 
   <td class="xl64">router</td> 
   <td class="xl64">string</td> 
   <td class="xl64">FALSE</td> 
   <td class="xl64">Path definition that separated by spaces</td> 
  </tr>  
 </tbody>
</table>