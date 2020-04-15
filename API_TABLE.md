# Kubelet's API
<table>
  <tbody>
	<tr>
	  <th>Kubelet API </th>
	  <th align="center">HTTP request</th>
	  <th align="center">Description</th>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /stats
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET /stats<br>
GET /stats/summary <br>
GET /stats/summary?only_cpu_and_memory=true<br>
GET /stats/container <br>
GET /stats/{namespace}/{podName}/{uid}/{containerName} <br>
GET /stats/{podName}/{containerName} </code>
	  </td>
	  <td align="left" >Return the performance stats of node, pods and containers</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /metrics
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET /metrics<br>
GET /metrics/cadvisor<br>
GET /metrics/probes<br>
GET /metrics/resource/v1alpha1</pre>
	  </td>
	  <td align="left" >Return information about node CPU and memory usage</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /logs
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET /logs<br>
GET /logs/{subpath} </pre>
	  </td>
	  <td align="left" >Logs from the node</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /spec
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET /spec</pre>
	  </td>
	  <td align="left" >Cached MachineInfo returned by cadvisor</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /pods
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET /pods</pre>
	  </td>
	  <td align="left" >List of pods</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /healthz
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET /healthz<br>
GET /healthz/log <br>
GET /healthz/ping<br>
GET /healthz/syncloop </pre>
	  </td>
	  <td align="left" >Check the state of the node</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /configz
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET /configz</pre>
	  </td>
	  <td align="left" >Kubelet's configurations</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /containerLogs
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET /containerLogs/{podNamespace}/{podID}/{containerName}</pre>
	  </td>
	  <td align="left" >Container's logs</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /run
		 </code>
	  </td>
	  <td align="left">         
		  <pre>POST /run/{podNamespace}/{podID}/{containerName}<br>
POST /run/{podNamespace}/{podID}/{uid}/{containerName} <br>
* The body of the request: <code>"cmd={command}"</code><br>
Example: <code>"cmd=ls /"</code></pre>
	  </td>
	  <td align="left" >Run command inside a container</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /exec
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET  /exec/{podNamespace}/{podID}/{containerName}?command={command}/&input=1&output=1&tty=1<br>
POST /exec/{podNamespace}//{containerName}?command={command}/&input=1&output=1&tty=1<br>
GET  /exec/{podNamespace}/{podID}/{uid}/{containerName}?command={command}/&input=1&output=1&tty=1<br>
POST /exec/{podNamespace}/{podID}/{uid}/{containerName}?command={command}/&input=1&output=1&tty=1</pre>
	  </td>
	  <td align="left" >Run command inside a container with option for stream (interactive)</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /cri
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET /cri/exec/{valueFrom302}?cmd={command}</pre>
	  </td>
	  <td align="left" >Run commands inside a container through the Container Runtime Interface (CRI)</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /attach
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET  /attach/{podNamespace}/{podID}/{containerName}<br>
POST /attach/{podNamespace}//{containerName}<br>
GET  /attach/{podNamespace}/{podID}/{uid}/{containerName}<br>
POST /attach/{podNamespace}/{podID}/{uid}/{containerName}</pre>
	  </td>
	  <td align="left" >Attach to a container</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /portForward
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET  /portForward/{podNamespace}/{podID}/{containerName}<br>
POST /portForward/{podNamespace}//{containerName}<br>
GET  /portForward/{podNamespace}/{podID}/{uid}/{containerName}<br>
POST /portForward/{podNamespace}/{podID}/{uid}/{containerName}</pre>
	  </td>
	  <td align="left" >Port forwarding inside the contianer</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /runningpods
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET  /runningpods</pre>
	  </td>
	  <td align="left" >List all the running pods</td>
	</tr>
	<tr>
	  <td>
		 <code class="rich-diff-level-one">
		 /debug
		 </code>
	  </td>
	  <td align="left">         
		  <pre>GET  /debug/pprof/{profile}<br>
GET /debug/flags/v<br>
PUT /debug/flags/v (body: {integer})</pre>
	  </td>
	  <td align="left" >List all the running pods</td>
	</tr>
  </tbody>
</table>