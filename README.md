# فن الحرب

## Fan al-Harb 
[![wercker status](https://app.wercker.com/status/35accf74192150e5425107b8362e69fa/s/master "wercker status")](https://app.wercker.com/project/byKey/35accf74192150e5425107b8362e69fa)

Load test the Google Kubernetes Engine Basic Cluster Kubernetes Horizontal Pod Autoscaler using a simple hello-world gohttp server. Vegeta will be used to generate and report the load, while jaggr and jplot will be used to visualize the report in realtime. 

Example of load generation using vegejaggrplot (pronounced: vplot):

```
echo 'GET http://localhost:8080' | \
    vegeta attack -rate 5000 -workers 100 -duration 10m | vegeta dump | \
    jaggr @count=rps \
          hist\[100,200,300,400,500\]:code \
          p25,p50,p95:latency \
          sum:bytes_in \
          sum:bytes_out | \
    jplot rps+code.hist.100+code.hist.200+code.hist.300+code.hist.400+code.hist.500 \
          latency.p95+latency.p50+latency.p25 \
          bytes_in.sum+bytes_out.sum
```

![](https://github.com/rs/jplot/blob/master/doc/vegeta.gif)
