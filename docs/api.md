# API Reference

Packages:

- [etos.eiffel-community.github.io/v1alpha1](#etoseiffel-communitygithubiov1alpha1)
- [etos.eiffel-community.github.io/v1alpha2](#etoseiffel-communitygithubiov1alpha2)

# etos.eiffel-community.github.io/v1alpha1

Resource Types:

- [Cluster](#cluster)

- [EnvironmentRequest](#environmentrequest)

- [Environment](#environment)

- [Provider](#provider)

- [TestRun](#testrun)




## Cluster
<sup><sup>[↩ Parent](#etoseiffel-communitygithubiov1alpha1 )</sup></sup>






Cluster is the Schema for the clusters API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>etos.eiffel-community.github.io/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>Cluster</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#clusterspec">spec</a></b></td>
        <td>object</td>
        <td>
          ClusterSpec defines the desired state of Cluster<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterstatus">status</a></b></td>
        <td>object</td>
        <td>
          ClusterStatus defines the observed state of Cluster<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec
<sup><sup>[↩ Parent](#cluster)</sup></sup>



ClusterSpec defines the desired state of Cluster

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#clusterspecdatabase">database</a></b></td>
        <td>object</td>
        <td>
          Database describes the deployment of a database for ETOS.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#clusterspecetos">etos</a></b></td>
        <td>object</td>
        <td>
          ETOS describes the deployment of an ETOS cluster.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#clusterspeceventrepository">eventRepository</a></b></td>
        <td>object</td>
        <td>
          EventRepository describes the deployment of an event repository for ETOS.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#clusterspecmessagebus">messageBus</a></b></td>
        <td>object</td>
        <td>
          MessageBus describes the deployment of messagesbuses for ETOS.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#clusterspecopentelemetry">openTelemetry</a></b></td>
        <td>object</td>
        <td>
          OpenTelemetry describes a deployment of an opentelemetry collector for ETOS to use.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### Cluster.spec.database
<sup><sup>[↩ Parent](#clusterspec)</sup></sup>



Database describes the deployment of a database for ETOS.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>deploy</b></td>
        <td>boolean</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: true<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecdatabaseetcd">etcd</a></b></td>
        <td>object</td>
        <td>
          Etcd describes the deployment of an ETCD database. Ignored if Deploy is set to false.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.database.etcd
<sup><sup>[↩ Parent](#clusterspecdatabase)</sup></sup>



Etcd describes the deployment of an ETCD database. Ignored if Deploy is set to false.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>host</b></td>
        <td>string</td>
        <td>
          Host specifies the ETCD server hostname.<br/>
          <br/>
            <i>Default</i>: etcd-client<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>port</b></td>
        <td>string</td>
        <td>
          Port specifies the ETCD port number.<br/>
          <br/>
            <i>Default</i>: 2379<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecdatabaseetcdresources">resources</a></b></td>
        <td>object</td>
        <td>
          Resources describes compute resource requirements per etcd pod which are three in a cluster.<br/>
          <br/>
            <i>Default</i>: map[limits:map[cpu:300m memory:768Mi] requests:map[cpu:300m memory:768Mi]]<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.database.etcd.resources
<sup><sup>[↩ Parent](#clusterspecdatabaseetcd)</sup></sup>



Resources describes compute resource requirements per etcd pod which are three in a cluster.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#clusterspecdatabaseetcdresourceslimits">limits</a></b></td>
        <td>object</td>
        <td>
          Limits describes the maximum amount of compute resources allowed.<br/>
          <br/>
            <i>Default</i>: map[cpu:300m memory:768Mi]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecdatabaseetcdresourcesrequests">requests</a></b></td>
        <td>object</td>
        <td>
          Requests describes the minimum amount of compute resources required.<br/>
          <br/>
            <i>Default</i>: map[cpu:300m memory:768Mi]<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.database.etcd.resources.limits
<sup><sup>[↩ Parent](#clusterspecdatabaseetcdresources)</sup></sup>



Limits describes the maximum amount of compute resources allowed.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>cpu</b></td>
        <td>string</td>
        <td>
          CPU resource.<br/>
          <br/>
            <i>Default</i>: 300m<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>memory</b></td>
        <td>string</td>
        <td>
          Memory resource.<br/>
          <br/>
            <i>Default</i>: 768Mi<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.database.etcd.resources.requests
<sup><sup>[↩ Parent](#clusterspecdatabaseetcdresources)</sup></sup>



Requests describes the minimum amount of compute resources required.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>cpu</b></td>
        <td>string</td>
        <td>
          CPU resource.<br/>
          <br/>
            <i>Default</i>: 300m<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>memory</b></td>
        <td>string</td>
        <td>
          Memory resource.<br/>
          <br/>
            <i>Default</i>: 768Mi<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos
<sup><sup>[↩ Parent](#clusterspec)</sup></sup>



ETOS describes the deployment of an ETOS cluster.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#clusterspecetosapi">api</a></b></td>
        <td>object</td>
        <td>
          ETOSAPI describes the deployment of the ETOS API.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecetosconfig">config</a></b></td>
        <td>object</td>
        <td>
          ETOSConfig describes a common configuration for ETOS.<br/>
          <br/>
            <i>Default</i>: map[encryptionKey:map[value:]]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecetosenvironmentprovider">environmentProvider</a></b></td>
        <td>object</td>
        <td>
          ETOSEnvironmentProvider describes the deployment of an ETOS environment provider.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecetosingress">ingress</a></b></td>
        <td>object</td>
        <td>
          Ingress configuration.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecetoslogarea">logArea</a></b></td>
        <td>object</td>
        <td>
          ETOSLogArea describes th deployment of an ETOS log area API.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecetossse">sse</a></b></td>
        <td>object</td>
        <td>
          ETOSSSE describes th deployment of an ETOS Server Sent Events API.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecetossuiterunner">suiteRunner</a></b></td>
        <td>object</td>
        <td>
          ETOSSuiteRunner describes the deployment of an ETOS suite runner.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecetossuitestarter">suiteStarter</a></b></td>
        <td>object</td>
        <td>
          ETOSSuiteStarter describes the deployment of an ETOS suite starter.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecetostestrunner">testRunner</a></b></td>
        <td>object</td>
        <td>
          ETOSTestRunner describes the deployment of an ETOS test runner.<br/>
          <br/>
            <i>Default</i>: map[version:3.7.5]<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.api
<sup><sup>[↩ Parent](#clusterspecetos)</sup></sup>



ETOSAPI describes the deployment of the ETOS API.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>executionSpaceProviderSecret</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: ghcr.io/eiffel-community/etos-api:1a5937dc<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>iutProviderSecret</b></td>
        <td>string</td>
        <td>
          The provider secrets are necessary in order deploy and run ETOS without using the
kubernetes controller.
They can be removed from here when the suite starter is no longer in use.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>logAreaProviderSecret</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>replicas</b></td>
        <td>integer</td>
        <td>
          <br/>
          <br/>
            <i>Format</i>: int32<br/>
            <i>Default</i>: 1<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.config
<sup><sup>[↩ Parent](#clusterspecetos)</sup></sup>



ETOSConfig describes a common configuration for ETOS.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#clusterspecetosconfigencryptionkey">encryptionKey</a></b></td>
        <td>object</td>
        <td>
          Var describes either a string value or a value from a VarSource.<br/>
          <br/>
            <i>Default</i>: map[value:]<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>dev</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: true<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>environmentTimeout</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 3600<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>etosApiURL</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>etosEventRepositoryURL</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>eventDataTimeout</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 60<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>routingKeyTag</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: etos<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>source</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: ETOS<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>testSuiteTimeout</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 10<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecetosconfigtestrunretention">testrunRetention</a></b></td>
        <td>object</td>
        <td>
          Retention describes the failure and success retentions for testruns.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>timezone</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.config.encryptionKey
<sup><sup>[↩ Parent](#clusterspecetosconfig)</sup></sup>



Var describes either a string value or a value from a VarSource.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecetosconfigencryptionkeyvaluefrom">valueFrom</a></b></td>
        <td>object</td>
        <td>
          VarSource describes a value from either a secretmap or configmap.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.config.encryptionKey.valueFrom
<sup><sup>[↩ Parent](#clusterspecetosconfigencryptionkey)</sup></sup>



VarSource describes a value from either a secretmap or configmap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#clusterspecetosconfigencryptionkeyvaluefromconfigmapkeyref">configMapKeyRef</a></b></td>
        <td>object</td>
        <td>
          Selects a key from a ConfigMap.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecetosconfigencryptionkeyvaluefromsecretkeyref">secretKeyRef</a></b></td>
        <td>object</td>
        <td>
          SecretKeySelector selects a key of a Secret.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.config.encryptionKey.valueFrom.configMapKeyRef
<sup><sup>[↩ Parent](#clusterspecetosconfigencryptionkeyvaluefrom)</sup></sup>



Selects a key from a ConfigMap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key to select.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the ConfigMap or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.config.encryptionKey.valueFrom.secretKeyRef
<sup><sup>[↩ Parent](#clusterspecetosconfigencryptionkeyvaluefrom)</sup></sup>



SecretKeySelector selects a key of a Secret.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key of the secret to select from.  Must be a valid secret key.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the Secret or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.config.testrunRetention
<sup><sup>[↩ Parent](#clusterspecetosconfig)</sup></sup>



Retention describes the failure and success retentions for testruns.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>failure</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>success</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.environmentProvider
<sup><sup>[↩ Parent](#clusterspecetos)</sup></sup>



ETOSEnvironmentProvider describes the deployment of an ETOS environment provider.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: ghcr.io/eiffel-community/etos-environment-provider:1c4413d5<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.ingress
<sup><sup>[↩ Parent](#clusterspecetos)</sup></sup>



Ingress configuration.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>annotations</b></td>
        <td>map[string]string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>enabled</b></td>
        <td>boolean</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>host</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>ingressClass</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.logArea
<sup><sup>[↩ Parent](#clusterspecetos)</sup></sup>



ETOSLogArea describes th deployment of an ETOS log area API.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: ghcr.io/eiffel-community/etos-logarea:1a5937dc<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.sse
<sup><sup>[↩ Parent](#clusterspecetos)</sup></sup>



ETOSSSE describes th deployment of an ETOS Server Sent Events API.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: ghcr.io/eiffel-community/etos-sse:1a5937dc<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.suiteRunner
<sup><sup>[↩ Parent](#clusterspecetos)</sup></sup>



ETOSSuiteRunner describes the deployment of an ETOS suite runner.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: ghcr.io/eiffel-community/etos-suite-runner:77de800f<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecetossuiterunnerloglistener">logListener</a></b></td>
        <td>object</td>
        <td>
          ETOSLogListener describes the deployment of an ETOS log listener.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.suiteRunner.logListener
<sup><sup>[↩ Parent](#clusterspecetossuiterunner)</sup></sup>



ETOSLogListener describes the deployment of an ETOS log listener.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>etosQueueName</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: etos-*-temp<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>etosQueueParams</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: ghcr.io/eiffel-community/etos-log-listener:77de800f<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.suiteStarter
<sup><sup>[↩ Parent](#clusterspecetos)</sup></sup>



ETOSSuiteStarter describes the deployment of an ETOS suite starter.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#clusterspecetossuitestarterconfig">config</a></b></td>
        <td>object</td>
        <td>
          ETOSSuiteStarterConfig describes the configuration required for a suite starter.
This is separate from the ETOSConfig as we want to remove this in the future when the suite
starter is no longer in use.<br/>
          <br/>
            <i>Default</i>: map[gracePeriod:300 ttl:3600]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>eiffelQueueName</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: etos-suite-starter<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>eiffelQueueParams</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: ghcr.io/eiffel-community/etos-suite-starter:6b1146e1<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>replicas</b></td>
        <td>integer</td>
        <td>
          <br/>
          <br/>
            <i>Format</i>: int32<br/>
            <i>Default</i>: 1<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>suiteRunnerTemplateSecretName</b></td>
        <td>string</td>
        <td>
          Provide a custom suite runner template.<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.suiteStarter.config
<sup><sup>[↩ Parent](#clusterspecetossuitestarter)</sup></sup>



ETOSSuiteStarterConfig describes the configuration required for a suite starter.
This is separate from the ETOSConfig as we want to remove this in the future when the suite
starter is no longer in use.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>gracePeriod</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 300<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>sidecarImage</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>ttl</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 3600<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.etos.testRunner
<sup><sup>[↩ Parent](#clusterspecetos)</sup></sup>



ETOSTestRunner describes the deployment of an ETOS test runner.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>version</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.eventRepository
<sup><sup>[↩ Parent](#clusterspec)</sup></sup>



EventRepository describes the deployment of an event repository for ETOS.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#clusterspeceventrepositoryapi">api</a></b></td>
        <td>object</td>
        <td>
          We do not build the GraphQL API automatically nor publish it remotely.
This will need to be provided to work.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>deploy</b></td>
        <td>boolean</td>
        <td>
          Deploy a local event repository for a cluster.<br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>eiffelQueueName</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: etos<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>eiffelQueueParams</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>host</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspeceventrepositoryingress">ingress</a></b></td>
        <td>object</td>
        <td>
          Ingress configuration.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspeceventrepositorymongo">mongo</a></b></td>
        <td>object</td>
        <td>
          MongoDB describes the deployment of a MongoDB.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspeceventrepositorystorage">storage</a></b></td>
        <td>object</td>
        <td>
          We do not build the GraphQL API automatically nor publish it remotely.
This will need to be provided to work.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.eventRepository.api
<sup><sup>[↩ Parent](#clusterspeceventrepository)</sup></sup>



We do not build the GraphQL API automatically nor publish it remotely.
This will need to be provided to work.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: ghcr.io/eiffel-community/eiffel-graphql-api:latest<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.eventRepository.ingress
<sup><sup>[↩ Parent](#clusterspeceventrepository)</sup></sup>



Ingress configuration.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>annotations</b></td>
        <td>map[string]string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>enabled</b></td>
        <td>boolean</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>host</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>ingressClass</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.eventRepository.mongo
<sup><sup>[↩ Parent](#clusterspeceventrepository)</sup></sup>



MongoDB describes the deployment of a MongoDB.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>deploy</b></td>
        <td>boolean</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspeceventrepositorymongouri">uri</a></b></td>
        <td>object</td>
        <td>
          Ignored if deploy is true<br/>
          <br/>
            <i>Default</i>: map[value:mongodb://root:password@mongodb:27017/admin]<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.eventRepository.mongo.uri
<sup><sup>[↩ Parent](#clusterspeceventrepositorymongo)</sup></sup>



Ignored if deploy is true

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspeceventrepositorymongourivaluefrom">valueFrom</a></b></td>
        <td>object</td>
        <td>
          VarSource describes a value from either a secretmap or configmap.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.eventRepository.mongo.uri.valueFrom
<sup><sup>[↩ Parent](#clusterspeceventrepositorymongouri)</sup></sup>



VarSource describes a value from either a secretmap or configmap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#clusterspeceventrepositorymongourivaluefromconfigmapkeyref">configMapKeyRef</a></b></td>
        <td>object</td>
        <td>
          Selects a key from a ConfigMap.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspeceventrepositorymongourivaluefromsecretkeyref">secretKeyRef</a></b></td>
        <td>object</td>
        <td>
          SecretKeySelector selects a key of a Secret.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.eventRepository.mongo.uri.valueFrom.configMapKeyRef
<sup><sup>[↩ Parent](#clusterspeceventrepositorymongourivaluefrom)</sup></sup>



Selects a key from a ConfigMap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key to select.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the ConfigMap or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.eventRepository.mongo.uri.valueFrom.secretKeyRef
<sup><sup>[↩ Parent](#clusterspeceventrepositorymongourivaluefrom)</sup></sup>



SecretKeySelector selects a key of a Secret.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key of the secret to select from.  Must be a valid secret key.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the Secret or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.eventRepository.storage
<sup><sup>[↩ Parent](#clusterspeceventrepository)</sup></sup>



We do not build the GraphQL API automatically nor publish it remotely.
This will need to be provided to work.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: ghcr.io/eiffel-community/eiffel-graphql-storage:latest<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.messageBus
<sup><sup>[↩ Parent](#clusterspec)</sup></sup>



MessageBus describes the deployment of messagesbuses for ETOS.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#clusterspecmessagebuseiffel">eiffel</a></b></td>
        <td>object</td>
        <td>
          RabbitMQ configuration.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecmessagebuslogs">logs</a></b></td>
        <td>object</td>
        <td>
          RabbitMQ configuration.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.messageBus.eiffel
<sup><sup>[↩ Parent](#clusterspecmessagebus)</sup></sup>



RabbitMQ configuration.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>deploy</b></td>
        <td>boolean</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>exchange</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: amq.topic<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>host</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: rabbitmq<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecmessagebuseiffelpassword">password</a></b></td>
        <td>object</td>
        <td>
          Var describes either a string value or a value from a VarSource.<br/>
          <br/>
            <i>Default</i>: map[value:guest]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>port</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 5672<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>username</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: guest<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>vhost</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: /<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.messageBus.eiffel.password
<sup><sup>[↩ Parent](#clusterspecmessagebuseiffel)</sup></sup>



Var describes either a string value or a value from a VarSource.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecmessagebuseiffelpasswordvaluefrom">valueFrom</a></b></td>
        <td>object</td>
        <td>
          VarSource describes a value from either a secretmap or configmap.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.messageBus.eiffel.password.valueFrom
<sup><sup>[↩ Parent](#clusterspecmessagebuseiffelpassword)</sup></sup>



VarSource describes a value from either a secretmap or configmap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#clusterspecmessagebuseiffelpasswordvaluefromconfigmapkeyref">configMapKeyRef</a></b></td>
        <td>object</td>
        <td>
          Selects a key from a ConfigMap.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecmessagebuseiffelpasswordvaluefromsecretkeyref">secretKeyRef</a></b></td>
        <td>object</td>
        <td>
          SecretKeySelector selects a key of a Secret.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.messageBus.eiffel.password.valueFrom.configMapKeyRef
<sup><sup>[↩ Parent](#clusterspecmessagebuseiffelpasswordvaluefrom)</sup></sup>



Selects a key from a ConfigMap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key to select.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the ConfigMap or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.messageBus.eiffel.password.valueFrom.secretKeyRef
<sup><sup>[↩ Parent](#clusterspecmessagebuseiffelpasswordvaluefrom)</sup></sup>



SecretKeySelector selects a key of a Secret.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key of the secret to select from.  Must be a valid secret key.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the Secret or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.messageBus.logs
<sup><sup>[↩ Parent](#clusterspecmessagebus)</sup></sup>



RabbitMQ configuration.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>deploy</b></td>
        <td>boolean</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>exchange</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: amq.topic<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>host</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: rabbitmq<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecmessagebuslogspassword">password</a></b></td>
        <td>object</td>
        <td>
          Var describes either a string value or a value from a VarSource.<br/>
          <br/>
            <i>Default</i>: map[value:guest]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>port</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 5672<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>username</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: guest<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>vhost</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: /<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.messageBus.logs.password
<sup><sup>[↩ Parent](#clusterspecmessagebuslogs)</sup></sup>



Var describes either a string value or a value from a VarSource.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecmessagebuslogspasswordvaluefrom">valueFrom</a></b></td>
        <td>object</td>
        <td>
          VarSource describes a value from either a secretmap or configmap.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.messageBus.logs.password.valueFrom
<sup><sup>[↩ Parent](#clusterspecmessagebuslogspassword)</sup></sup>



VarSource describes a value from either a secretmap or configmap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#clusterspecmessagebuslogspasswordvaluefromconfigmapkeyref">configMapKeyRef</a></b></td>
        <td>object</td>
        <td>
          Selects a key from a ConfigMap.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#clusterspecmessagebuslogspasswordvaluefromsecretkeyref">secretKeyRef</a></b></td>
        <td>object</td>
        <td>
          SecretKeySelector selects a key of a Secret.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.messageBus.logs.password.valueFrom.configMapKeyRef
<sup><sup>[↩ Parent](#clusterspecmessagebuslogspasswordvaluefrom)</sup></sup>



Selects a key from a ConfigMap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key to select.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the ConfigMap or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.messageBus.logs.password.valueFrom.secretKeyRef
<sup><sup>[↩ Parent](#clusterspecmessagebuslogspasswordvaluefrom)</sup></sup>



SecretKeySelector selects a key of a Secret.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key of the secret to select from.  Must be a valid secret key.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the Secret or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.spec.openTelemetry
<sup><sup>[↩ Parent](#clusterspec)</sup></sup>



OpenTelemetry describes a deployment of an opentelemetry collector for ETOS to use.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>enabled</b></td>
        <td>boolean</td>
        <td>
          Enable opentelemetry support, adding the environment variables to services.<br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td>
          Sets the OTEL_EXPORTER_OTLP_ENDPOINT environment variable<br/>
          <br/>
            <i>Default</i>: http://localhost:4317<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>insecure</b></td>
        <td>string</td>
        <td>
          Sets the OTEL_EXPORTER_OTLP_INSECURE environment variable<br/>
          <br/>
            <i>Default</i>: true<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.status
<sup><sup>[↩ Parent](#cluster)</sup></sup>



ClusterStatus defines the observed state of Cluster

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#clusterstatusconditionsindex">conditions</a></b></td>
        <td>[]object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Cluster.status.conditions[index]
<sup><sup>[↩ Parent](#clusterstatus)</sup></sup>



Condition contains details for one aspect of the current state of this API Resource.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>lastTransitionTime</b></td>
        <td>string</td>
        <td>
          lastTransitionTime is the last time the condition transitioned from one status to another.
This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.<br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>message</b></td>
        <td>string</td>
        <td>
          message is a human readable message indicating details about the transition.
This may be an empty string.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>reason</b></td>
        <td>string</td>
        <td>
          reason contains a programmatic identifier indicating the reason for the condition's last transition.
Producers of specific condition types may define expected values and meanings for this field,
and whether the values are considered a guaranteed API.
The value should be a CamelCase string.
This field may not be empty.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>status</b></td>
        <td>enum</td>
        <td>
          status of the condition, one of True, False, Unknown.<br/>
          <br/>
            <i>Enum</i>: True, False, Unknown<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          type of condition in CamelCase or in foo.example.com/CamelCase.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>observedGeneration</b></td>
        <td>integer</td>
        <td>
          observedGeneration represents the .metadata.generation that the condition was set based upon.
For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date
with respect to the current state of the instance.<br/>
          <br/>
            <i>Format</i>: int64<br/>
            <i>Minimum</i>: 0<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>

## EnvironmentRequest
<sup><sup>[↩ Parent](#etoseiffel-communitygithubiov1alpha1 )</sup></sup>






EnvironmentRequest is the Schema for the environmentrequests API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>etos.eiffel-community.github.io/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>EnvironmentRequest</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspec">spec</a></b></td>
        <td>object</td>
        <td>
          EnvironmentRequestSpec defines the desired state of EnvironmentRequest<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentrequeststatus">status</a></b></td>
        <td>object</td>
        <td>
          EnvironmentRequestStatus defines the observed state of EnvironmentRequest<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec
<sup><sup>[↩ Parent](#environmentrequest)</sup></sup>



EnvironmentRequestSpec defines the desired state of EnvironmentRequest

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>maximumAmount</b></td>
        <td>integer</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>minimumAmount</b></td>
        <td>integer</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecproviders">providers</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecsplitter">splitter</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecconfig">Config</a></b></td>
        <td>object</td>
        <td>
          EnvironmentProviderJobConfig defines parameters required by environment provider job<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>artifact</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>dataset</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          ID is the ID for the environments generated. Will be generated if nil. The ID is a UUID, any version, and regex matches that.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>identifier</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>identity</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>serviceaccountname</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.providers
<sup><sup>[↩ Parent](#environmentrequestspec)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#environmentrequestspecprovidersexecutionspace">executionSpace</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecprovidersiut">iut</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecproviderslogarea">logArea</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.providers.executionSpace
<sup><sup>[↩ Parent](#environmentrequestspecproviders)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>testRunner</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.providers.iut
<sup><sup>[↩ Parent](#environmentrequestspecproviders)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.providers.logArea
<sup><sup>[↩ Parent](#environmentrequestspecproviders)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.splitter
<sup><sup>[↩ Parent](#environmentrequestspec)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#environmentrequestspecsplittertestsindex">tests</a></b></td>
        <td>[]object</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.splitter.tests[index]
<sup><sup>[↩ Parent](#environmentrequestspecsplitter)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>environment</b></td>
        <td>object</td>
        <td>
          TestEnvironment to run tests within.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecsplittertestsindexexecution">execution</a></b></td>
        <td>object</td>
        <td>
          Execution describes how to execute a testCase.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecsplittertestsindextestcase">testCase</a></b></td>
        <td>object</td>
        <td>
          TestCase metadata.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.splitter.tests[index].execution
<sup><sup>[↩ Parent](#environmentrequestspecsplittertestsindex)</sup></sup>



Execution describes how to execute a testCase.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>checkout</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>command</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>environment</b></td>
        <td>map[string]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>parameters</b></td>
        <td>map[string]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>testRunner</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>execute</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.splitter.tests[index].testCase
<sup><sup>[↩ Parent](#environmentrequestspecsplittertestsindex)</sup></sup>



TestCase metadata.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>tracker</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>uri</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>version</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config
<sup><sup>[↩ Parent](#environmentrequestspec)</sup></sup>



EnvironmentProviderJobConfig defines parameters required by environment provider job

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#environmentrequestspecconfigeiffelmessagebus">eiffelMessageBus</a></b></td>
        <td>object</td>
        <td>
          RabbitMQ configuration.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecconfigencryptionkeysecretref">encryptionKeySecretRef</a></b></td>
        <td>object</td>
        <td>
          Var describes either a string value or a value from a VarSource.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>environmentProviderEventDataTimeout</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>environmentProviderImage</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>environmentProviderImagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>environmentProviderServiceAccount</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>environmentProviderTestSuiteTimeout</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>etcdHost</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>etcdPort</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>etosApi</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecconfigetosmessagebus">etosMessageBus</a></b></td>
        <td>object</td>
        <td>
          RabbitMQ configuration.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>graphQlServer</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>routingKeyTag</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>testRunnerVersion</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>waitForTimeout</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.eiffelMessageBus
<sup><sup>[↩ Parent](#environmentrequestspecconfig)</sup></sup>



RabbitMQ configuration.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>deploy</b></td>
        <td>boolean</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>exchange</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: amq.topic<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>host</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: rabbitmq<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecconfigeiffelmessagebuspassword">password</a></b></td>
        <td>object</td>
        <td>
          Var describes either a string value or a value from a VarSource.<br/>
          <br/>
            <i>Default</i>: map[value:guest]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>port</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 5672<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>username</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: guest<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>vhost</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: /<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.eiffelMessageBus.password
<sup><sup>[↩ Parent](#environmentrequestspecconfigeiffelmessagebus)</sup></sup>



Var describes either a string value or a value from a VarSource.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecconfigeiffelmessagebuspasswordvaluefrom">valueFrom</a></b></td>
        <td>object</td>
        <td>
          VarSource describes a value from either a secretmap or configmap.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.eiffelMessageBus.password.valueFrom
<sup><sup>[↩ Parent](#environmentrequestspecconfigeiffelmessagebuspassword)</sup></sup>



VarSource describes a value from either a secretmap or configmap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#environmentrequestspecconfigeiffelmessagebuspasswordvaluefromconfigmapkeyref">configMapKeyRef</a></b></td>
        <td>object</td>
        <td>
          Selects a key from a ConfigMap.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecconfigeiffelmessagebuspasswordvaluefromsecretkeyref">secretKeyRef</a></b></td>
        <td>object</td>
        <td>
          SecretKeySelector selects a key of a Secret.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.eiffelMessageBus.password.valueFrom.configMapKeyRef
<sup><sup>[↩ Parent](#environmentrequestspecconfigeiffelmessagebuspasswordvaluefrom)</sup></sup>



Selects a key from a ConfigMap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key to select.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the ConfigMap or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.eiffelMessageBus.password.valueFrom.secretKeyRef
<sup><sup>[↩ Parent](#environmentrequestspecconfigeiffelmessagebuspasswordvaluefrom)</sup></sup>



SecretKeySelector selects a key of a Secret.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key of the secret to select from.  Must be a valid secret key.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the Secret or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.encryptionKeySecretRef
<sup><sup>[↩ Parent](#environmentrequestspecconfig)</sup></sup>



Var describes either a string value or a value from a VarSource.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecconfigencryptionkeysecretrefvaluefrom">valueFrom</a></b></td>
        <td>object</td>
        <td>
          VarSource describes a value from either a secretmap or configmap.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.encryptionKeySecretRef.valueFrom
<sup><sup>[↩ Parent](#environmentrequestspecconfigencryptionkeysecretref)</sup></sup>



VarSource describes a value from either a secretmap or configmap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#environmentrequestspecconfigencryptionkeysecretrefvaluefromconfigmapkeyref">configMapKeyRef</a></b></td>
        <td>object</td>
        <td>
          Selects a key from a ConfigMap.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecconfigencryptionkeysecretrefvaluefromsecretkeyref">secretKeyRef</a></b></td>
        <td>object</td>
        <td>
          SecretKeySelector selects a key of a Secret.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.encryptionKeySecretRef.valueFrom.configMapKeyRef
<sup><sup>[↩ Parent](#environmentrequestspecconfigencryptionkeysecretrefvaluefrom)</sup></sup>



Selects a key from a ConfigMap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key to select.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the ConfigMap or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.encryptionKeySecretRef.valueFrom.secretKeyRef
<sup><sup>[↩ Parent](#environmentrequestspecconfigencryptionkeysecretrefvaluefrom)</sup></sup>



SecretKeySelector selects a key of a Secret.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key of the secret to select from.  Must be a valid secret key.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the Secret or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.etosMessageBus
<sup><sup>[↩ Parent](#environmentrequestspecconfig)</sup></sup>



RabbitMQ configuration.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>deploy</b></td>
        <td>boolean</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>exchange</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: amq.topic<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>host</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: rabbitmq<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecconfigetosmessagebuspassword">password</a></b></td>
        <td>object</td>
        <td>
          Var describes either a string value or a value from a VarSource.<br/>
          <br/>
            <i>Default</i>: map[value:guest]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>port</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 5672<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>username</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: guest<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>vhost</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: /<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.etosMessageBus.password
<sup><sup>[↩ Parent](#environmentrequestspecconfigetosmessagebus)</sup></sup>



Var describes either a string value or a value from a VarSource.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecconfigetosmessagebuspasswordvaluefrom">valueFrom</a></b></td>
        <td>object</td>
        <td>
          VarSource describes a value from either a secretmap or configmap.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.etosMessageBus.password.valueFrom
<sup><sup>[↩ Parent](#environmentrequestspecconfigetosmessagebuspassword)</sup></sup>



VarSource describes a value from either a secretmap or configmap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#environmentrequestspecconfigetosmessagebuspasswordvaluefromconfigmapkeyref">configMapKeyRef</a></b></td>
        <td>object</td>
        <td>
          Selects a key from a ConfigMap.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentrequestspecconfigetosmessagebuspasswordvaluefromsecretkeyref">secretKeyRef</a></b></td>
        <td>object</td>
        <td>
          SecretKeySelector selects a key of a Secret.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.etosMessageBus.password.valueFrom.configMapKeyRef
<sup><sup>[↩ Parent](#environmentrequestspecconfigetosmessagebuspasswordvaluefrom)</sup></sup>



Selects a key from a ConfigMap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key to select.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the ConfigMap or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.spec.Config.etosMessageBus.password.valueFrom.secretKeyRef
<sup><sup>[↩ Parent](#environmentrequestspecconfigetosmessagebuspasswordvaluefrom)</sup></sup>



SecretKeySelector selects a key of a Secret.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key of the secret to select from.  Must be a valid secret key.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the Secret or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.status
<sup><sup>[↩ Parent](#environmentrequest)</sup></sup>



EnvironmentRequestStatus defines the observed state of EnvironmentRequest

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>completionTime</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentrequeststatusconditionsindex">conditions</a></b></td>
        <td>[]object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentrequeststatusenvironmentprovidersindex">environmentProviders</a></b></td>
        <td>[]object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>startTime</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.status.conditions[index]
<sup><sup>[↩ Parent](#environmentrequeststatus)</sup></sup>



Condition contains details for one aspect of the current state of this API Resource.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>lastTransitionTime</b></td>
        <td>string</td>
        <td>
          lastTransitionTime is the last time the condition transitioned from one status to another.
This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.<br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>message</b></td>
        <td>string</td>
        <td>
          message is a human readable message indicating details about the transition.
This may be an empty string.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>reason</b></td>
        <td>string</td>
        <td>
          reason contains a programmatic identifier indicating the reason for the condition's last transition.
Producers of specific condition types may define expected values and meanings for this field,
and whether the values are considered a guaranteed API.
The value should be a CamelCase string.
This field may not be empty.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>status</b></td>
        <td>enum</td>
        <td>
          status of the condition, one of True, False, Unknown.<br/>
          <br/>
            <i>Enum</i>: True, False, Unknown<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          type of condition in CamelCase or in foo.example.com/CamelCase.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>observedGeneration</b></td>
        <td>integer</td>
        <td>
          observedGeneration represents the .metadata.generation that the condition was set based upon.
For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date
with respect to the current state of the instance.<br/>
          <br/>
            <i>Format</i>: int64<br/>
            <i>Minimum</i>: 0<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EnvironmentRequest.status.environmentProviders[index]
<sup><sup>[↩ Parent](#environmentrequeststatus)</sup></sup>



ObjectReference contains enough information to let you inspect or modify the referred object.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>apiVersion</b></td>
        <td>string</td>
        <td>
          API version of the referent.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>fieldPath</b></td>
        <td>string</td>
        <td>
          If referring to a piece of an object instead of an entire object, this string
should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
For example, if the object reference is to a container within a pod, this would take on a value like:
"spec.containers{name}" (where "name" refers to the name of the container that triggered
the event) or if no container name is specified "spec.containers[2]" (container with
index 2 in this pod). This syntax is chosen only to have some well-defined way of
referencing a part of an object.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>kind</b></td>
        <td>string</td>
        <td>
          Kind of the referent.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>namespace</b></td>
        <td>string</td>
        <td>
          Namespace of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>resourceVersion</b></td>
        <td>string</td>
        <td>
          Specific resourceVersion to which this reference is made, if any.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>uid</b></td>
        <td>string</td>
        <td>
          UID of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>

## Environment
<sup><sup>[↩ Parent](#etoseiffel-communitygithubiov1alpha1 )</sup></sup>






Environment is the Schema for the environments API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>etos.eiffel-community.github.io/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>Environment</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#environmentspec">spec</a></b></td>
        <td>object</td>
        <td>
          EnvironmentSpec defines the desired state of Environment<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentstatus">status</a></b></td>
        <td>object</td>
        <td>
          EnvironmentStatus defines the observed state of Environment<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Environment.spec
<sup><sup>[↩ Parent](#environment)</sup></sup>



EnvironmentSpec defines the desired state of Environment

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>artifact</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>context</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>executor</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>iut</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>log_area</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>priority</b></td>
        <td>integer</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#environmentspecrecipesindex">recipes</a></b></td>
        <td>[]object</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>sub_suite_id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>suite_id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>test_runner</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>test_suite_started_id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### Environment.spec.recipes[index]
<sup><sup>[↩ Parent](#environmentspec)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>environment</b></td>
        <td>object</td>
        <td>
          TestEnvironment to run tests within.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#environmentspecrecipesindexexecution">execution</a></b></td>
        <td>object</td>
        <td>
          Execution describes how to execute a testCase.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#environmentspecrecipesindextestcase">testCase</a></b></td>
        <td>object</td>
        <td>
          TestCase metadata.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### Environment.spec.recipes[index].execution
<sup><sup>[↩ Parent](#environmentspecrecipesindex)</sup></sup>



Execution describes how to execute a testCase.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>checkout</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>command</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>environment</b></td>
        <td>map[string]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>parameters</b></td>
        <td>map[string]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>testRunner</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>execute</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Environment.spec.recipes[index].testCase
<sup><sup>[↩ Parent](#environmentspecrecipesindex)</sup></sup>



TestCase metadata.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>tracker</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>uri</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>version</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Environment.status
<sup><sup>[↩ Parent](#environment)</sup></sup>



EnvironmentStatus defines the observed state of Environment

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>completionTime</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentstatusconditionsindex">conditions</a></b></td>
        <td>[]object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#environmentstatusenvironmentreleasersindex">environmentReleasers</a></b></td>
        <td>[]object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Environment.status.conditions[index]
<sup><sup>[↩ Parent](#environmentstatus)</sup></sup>



Condition contains details for one aspect of the current state of this API Resource.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>lastTransitionTime</b></td>
        <td>string</td>
        <td>
          lastTransitionTime is the last time the condition transitioned from one status to another.
This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.<br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>message</b></td>
        <td>string</td>
        <td>
          message is a human readable message indicating details about the transition.
This may be an empty string.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>reason</b></td>
        <td>string</td>
        <td>
          reason contains a programmatic identifier indicating the reason for the condition's last transition.
Producers of specific condition types may define expected values and meanings for this field,
and whether the values are considered a guaranteed API.
The value should be a CamelCase string.
This field may not be empty.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>status</b></td>
        <td>enum</td>
        <td>
          status of the condition, one of True, False, Unknown.<br/>
          <br/>
            <i>Enum</i>: True, False, Unknown<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          type of condition in CamelCase or in foo.example.com/CamelCase.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>observedGeneration</b></td>
        <td>integer</td>
        <td>
          observedGeneration represents the .metadata.generation that the condition was set based upon.
For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date
with respect to the current state of the instance.<br/>
          <br/>
            <i>Format</i>: int64<br/>
            <i>Minimum</i>: 0<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Environment.status.environmentReleasers[index]
<sup><sup>[↩ Parent](#environmentstatus)</sup></sup>



ObjectReference contains enough information to let you inspect or modify the referred object.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>apiVersion</b></td>
        <td>string</td>
        <td>
          API version of the referent.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>fieldPath</b></td>
        <td>string</td>
        <td>
          If referring to a piece of an object instead of an entire object, this string
should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
For example, if the object reference is to a container within a pod, this would take on a value like:
"spec.containers{name}" (where "name" refers to the name of the container that triggered
the event) or if no container name is specified "spec.containers[2]" (container with
index 2 in this pod). This syntax is chosen only to have some well-defined way of
referencing a part of an object.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>kind</b></td>
        <td>string</td>
        <td>
          Kind of the referent.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>namespace</b></td>
        <td>string</td>
        <td>
          Namespace of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>resourceVersion</b></td>
        <td>string</td>
        <td>
          Specific resourceVersion to which this reference is made, if any.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>uid</b></td>
        <td>string</td>
        <td>
          UID of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>

## Provider
<sup><sup>[↩ Parent](#etoseiffel-communitygithubiov1alpha1 )</sup></sup>






Provider is the Schema for the providers API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>etos.eiffel-community.github.io/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>Provider</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#providerspec">spec</a></b></td>
        <td>object</td>
        <td>
          ProviderSpec defines the desired state of Provider<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#providerstatus">status</a></b></td>
        <td>object</td>
        <td>
          ProviderStatus defines the observed state of Provider<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.spec
<sup><sup>[↩ Parent](#provider)</sup></sup>



ProviderSpec defines the desired state of Provider

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>type</b></td>
        <td>enum</td>
        <td>
          <br/>
          <br/>
            <i>Enum</i>: execution-space, iut, log-area<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#providerspechealthcheck">healthCheck</a></b></td>
        <td>object</td>
        <td>
          Healthcheck defines the health check endpoint and interval for providers.
The defaults of this should work most of the time.<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>host</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          Image describes the docker image to run when providing a resource.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#providerspecjsontas">jsontas</a></b></td>
        <td>object</td>
        <td>
          These are pointers so that they become nil in the Provider object in Kubernetes
and don't muddle up the yaml with empty data.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#providerspecjsontassource">jsontasSource</a></b></td>
        <td>object</td>
        <td>
          VarSource describes a value from either a secretmap or configmap.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.spec.healthCheck
<sup><sup>[↩ Parent](#providerspec)</sup></sup>



Healthcheck defines the health check endpoint and interval for providers.
The defaults of this should work most of the time.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: /v1alpha1/selftest/ping<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>intervalSeconds</b></td>
        <td>integer</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 30<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.spec.jsontas
<sup><sup>[↩ Parent](#providerspec)</sup></sup>



These are pointers so that they become nil in the Provider object in Kubernetes
and don't muddle up the yaml with empty data.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#providerspecjsontasexecution_space">execution_space</a></b></td>
        <td>object</td>
        <td>
          JSONTasExecutionSpace is the execution space provider definition for the JSONTas provider<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#providerspecjsontasiut">iut</a></b></td>
        <td>object</td>
        <td>
          These are pointers so that they become nil in the Provider object in Kubernetes
and don't muddle up the yaml with empty data.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#providerspecjsontaslog">log</a></b></td>
        <td>object</td>
        <td>
          JSONTasLogArea is the log area provider definition for the JSONTas provider<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.spec.jsontas.execution_space
<sup><sup>[↩ Parent](#providerspecjsontas)</sup></sup>



JSONTasExecutionSpace is the execution space provider definition for the JSONTas provider

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#providerspecjsontasexecution_spacelist">list</a></b></td>
        <td>object</td>
        <td>
          JSONTasList is the List command in the JSONTas provider.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>checkin</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>checkout</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.spec.jsontas.execution_space.list
<sup><sup>[↩ Parent](#providerspecjsontasexecution_space)</sup></sup>



JSONTasList is the List command in the JSONTas provider.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>available</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>possible</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### Provider.spec.jsontas.iut
<sup><sup>[↩ Parent](#providerspecjsontas)</sup></sup>



These are pointers so that they become nil in the Provider object in Kubernetes
and don't muddle up the yaml with empty data.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#providerspecjsontasiutlist">list</a></b></td>
        <td>object</td>
        <td>
          JSONTasList is the List command in the JSONTas provider.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>checkin</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>checkout</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#providerspecjsontasiutprepare">prepare</a></b></td>
        <td>object</td>
        <td>
          JSONTasIUTPrepare defines the preparation required for an IUT.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.spec.jsontas.iut.list
<sup><sup>[↩ Parent](#providerspecjsontasiut)</sup></sup>



JSONTasList is the List command in the JSONTas provider.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>available</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>possible</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### Provider.spec.jsontas.iut.prepare
<sup><sup>[↩ Parent](#providerspecjsontasiut)</sup></sup>



JSONTasIUTPrepare defines the preparation required for an IUT.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#providerspecjsontasiutpreparestages">stages</a></b></td>
        <td>object</td>
        <td>
          JSONTasIUTPrepareStages defines the preparation stages required for an IUT.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### Provider.spec.jsontas.iut.prepare.stages
<sup><sup>[↩ Parent](#providerspecjsontasiutprepare)</sup></sup>



JSONTasIUTPrepareStages defines the preparation stages required for an IUT.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#providerspecjsontasiutpreparestagesenvironment_provider">environment_provider</a></b></td>
        <td>object</td>
        <td>
          Underscore used in these due to backwards compatibility<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#providerspecjsontasiutpreparestagessuite_runner">suite_runner</a></b></td>
        <td>object</td>
        <td>
          Stage is the definition of a stage where to execute steps.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#providerspecjsontasiutpreparestagestest_runner">test_runner</a></b></td>
        <td>object</td>
        <td>
          Stage is the definition of a stage where to execute steps.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### Provider.spec.jsontas.iut.prepare.stages.environment_provider
<sup><sup>[↩ Parent](#providerspecjsontasiutpreparestages)</sup></sup>



Underscore used in these due to backwards compatibility

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>steps</b></td>
        <td>JSON</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.spec.jsontas.iut.prepare.stages.suite_runner
<sup><sup>[↩ Parent](#providerspecjsontasiutpreparestages)</sup></sup>



Stage is the definition of a stage where to execute steps.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>steps</b></td>
        <td>JSON</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.spec.jsontas.iut.prepare.stages.test_runner
<sup><sup>[↩ Parent](#providerspecjsontasiutpreparestages)</sup></sup>



Stage is the definition of a stage where to execute steps.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>steps</b></td>
        <td>JSON</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.spec.jsontas.log
<sup><sup>[↩ Parent](#providerspecjsontas)</sup></sup>



JSONTasLogArea is the log area provider definition for the JSONTas provider

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#providerspecjsontasloglist">list</a></b></td>
        <td>object</td>
        <td>
          JSONTasList is the List command in the JSONTas provider.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>checkin</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>checkout</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.spec.jsontas.log.list
<sup><sup>[↩ Parent](#providerspecjsontaslog)</sup></sup>



JSONTasList is the List command in the JSONTas provider.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>available</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>possible</b></td>
        <td>JSON</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### Provider.spec.jsontasSource
<sup><sup>[↩ Parent](#providerspec)</sup></sup>



VarSource describes a value from either a secretmap or configmap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#providerspecjsontassourceconfigmapkeyref">configMapKeyRef</a></b></td>
        <td>object</td>
        <td>
          Selects a key from a ConfigMap.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#providerspecjsontassourcesecretkeyref">secretKeyRef</a></b></td>
        <td>object</td>
        <td>
          SecretKeySelector selects a key of a Secret.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.spec.jsontasSource.configMapKeyRef
<sup><sup>[↩ Parent](#providerspecjsontassource)</sup></sup>



Selects a key from a ConfigMap.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key to select.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the ConfigMap or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.spec.jsontasSource.secretKeyRef
<sup><sup>[↩ Parent](#providerspecjsontassource)</sup></sup>



SecretKeySelector selects a key of a Secret.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The key of the secret to select from.  Must be a valid secret key.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
          <br/>
            <i>Default</i>: <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>optional</b></td>
        <td>boolean</td>
        <td>
          Specify whether the Secret or its key must be defined<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.status
<sup><sup>[↩ Parent](#provider)</sup></sup>



ProviderStatus defines the observed state of Provider

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#providerstatusconditionsindex">conditions</a></b></td>
        <td>[]object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>lastHealthCheckTime</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Provider.status.conditions[index]
<sup><sup>[↩ Parent](#providerstatus)</sup></sup>



Condition contains details for one aspect of the current state of this API Resource.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>lastTransitionTime</b></td>
        <td>string</td>
        <td>
          lastTransitionTime is the last time the condition transitioned from one status to another.
This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.<br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>message</b></td>
        <td>string</td>
        <td>
          message is a human readable message indicating details about the transition.
This may be an empty string.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>reason</b></td>
        <td>string</td>
        <td>
          reason contains a programmatic identifier indicating the reason for the condition's last transition.
Producers of specific condition types may define expected values and meanings for this field,
and whether the values are considered a guaranteed API.
The value should be a CamelCase string.
This field may not be empty.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>status</b></td>
        <td>enum</td>
        <td>
          status of the condition, one of True, False, Unknown.<br/>
          <br/>
            <i>Enum</i>: True, False, Unknown<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          type of condition in CamelCase or in foo.example.com/CamelCase.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>observedGeneration</b></td>
        <td>integer</td>
        <td>
          observedGeneration represents the .metadata.generation that the condition was set based upon.
For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date
with respect to the current state of the instance.<br/>
          <br/>
            <i>Format</i>: int64<br/>
            <i>Minimum</i>: 0<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>

## TestRun
<sup><sup>[↩ Parent](#etoseiffel-communitygithubiov1alpha1 )</sup></sup>






TestRun is the Schema for the testruns API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>etos.eiffel-community.github.io/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>TestRun</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#testrunspec">spec</a></b></td>
        <td>object</td>
        <td>
          TestRunSpec defines the desired state of TestRun<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#testrunstatus">status</a></b></td>
        <td>object</td>
        <td>
          TestRunStatus defines the observed state of TestRun<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### TestRun.spec
<sup><sup>[↩ Parent](#testrun)</sup></sup>



TestRunSpec defines the desired state of TestRun

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>artifact</b></td>
        <td>string</td>
        <td>
          Artifact is the ID of the software under test. The ID is a UUID, any version, and regex matches that.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>identity</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#testrunspecproviders">providers</a></b></td>
        <td>object</td>
        <td>
          Providers to use for test execution. These names must correspond to existing
Provider kinds in the namespace where a testrun is created.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#testrunspecsuitesindex">suites</a></b></td>
        <td>[]object</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>cluster</b></td>
        <td>string</td>
        <td>
          Name of the ETOS cluster to execute the testrun in.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#testrunspecenvironmentprovider">environmentProvider</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          ID is the test suite ID for this execution. Will be generated if nil. The ID is a UUID, any version, and regex matches that.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#testrunspecloglistener">logListener</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#testrunspecretention">retention</a></b></td>
        <td>object</td>
        <td>
          Retention describes the failure and success retentions for testruns.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#testrunspecsuiterunner">suiteRunner</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#testrunspectestrunner">testRunner</a></b></td>
        <td>object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### TestRun.spec.providers
<sup><sup>[↩ Parent](#testrunspec)</sup></sup>



Providers to use for test execution. These names must correspond to existing
Provider kinds in the namespace where a testrun is created.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>executionSpace</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>iut</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>logArea</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### TestRun.spec.suites[index]
<sup><sup>[↩ Parent](#testrunspec)</sup></sup>



Suite to execute.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>dataset</b></td>
        <td>JSON</td>
        <td>
          Dataset for this suite.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the test suite.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>priority</b></td>
        <td>integer</td>
        <td>
          Priority to execute the test suite.<br/>
          <br/>
            <i>Default</i>: 1<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#testrunspecsuitesindextestsindex">tests</a></b></td>
        <td>[]object</td>
        <td>
          Tests to execute as part of this suite.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### TestRun.spec.suites[index].tests[index]
<sup><sup>[↩ Parent](#testrunspecsuitesindex)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>environment</b></td>
        <td>object</td>
        <td>
          TestEnvironment to run tests within.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#testrunspecsuitesindextestsindexexecution">execution</a></b></td>
        <td>object</td>
        <td>
          Execution describes how to execute a testCase.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#testrunspecsuitesindextestsindextestcase">testCase</a></b></td>
        <td>object</td>
        <td>
          TestCase metadata.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### TestRun.spec.suites[index].tests[index].execution
<sup><sup>[↩ Parent](#testrunspecsuitesindextestsindex)</sup></sup>



Execution describes how to execute a testCase.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>checkout</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>command</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>environment</b></td>
        <td>map[string]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>parameters</b></td>
        <td>map[string]string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>testRunner</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>execute</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### TestRun.spec.suites[index].tests[index].testCase
<sup><sup>[↩ Parent](#testrunspecsuitesindextestsindex)</sup></sup>



TestCase metadata.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>tracker</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>uri</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>version</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### TestRun.spec.environmentProvider
<sup><sup>[↩ Parent](#testrunspec)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### TestRun.spec.logListener
<sup><sup>[↩ Parent](#testrunspec)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### TestRun.spec.retention
<sup><sup>[↩ Parent](#testrunspec)</sup></sup>



Retention describes the failure and success retentions for testruns.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>failure</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>success</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### TestRun.spec.suiteRunner
<sup><sup>[↩ Parent](#testrunspec)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>
          PullPolicy describes a policy for if/when to pull a container image<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### TestRun.spec.testRunner
<sup><sup>[↩ Parent](#testrunspec)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>version</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### TestRun.status
<sup><sup>[↩ Parent](#testrun)</sup></sup>



TestRunStatus defines the observed state of TestRun

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>completionTime</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#testrunstatusconditionsindex">conditions</a></b></td>
        <td>[]object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#testrunstatusenvironmentrequestsindex">environmentRequests</a></b></td>
        <td>[]object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>startTime</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#testrunstatussuiterunnersindex">suiteRunners</a></b></td>
        <td>[]object</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>verdict</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### TestRun.status.conditions[index]
<sup><sup>[↩ Parent](#testrunstatus)</sup></sup>



Condition contains details for one aspect of the current state of this API Resource.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>lastTransitionTime</b></td>
        <td>string</td>
        <td>
          lastTransitionTime is the last time the condition transitioned from one status to another.
This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.<br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>message</b></td>
        <td>string</td>
        <td>
          message is a human readable message indicating details about the transition.
This may be an empty string.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>reason</b></td>
        <td>string</td>
        <td>
          reason contains a programmatic identifier indicating the reason for the condition's last transition.
Producers of specific condition types may define expected values and meanings for this field,
and whether the values are considered a guaranteed API.
The value should be a CamelCase string.
This field may not be empty.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>status</b></td>
        <td>enum</td>
        <td>
          status of the condition, one of True, False, Unknown.<br/>
          <br/>
            <i>Enum</i>: True, False, Unknown<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          type of condition in CamelCase or in foo.example.com/CamelCase.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>observedGeneration</b></td>
        <td>integer</td>
        <td>
          observedGeneration represents the .metadata.generation that the condition was set based upon.
For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date
with respect to the current state of the instance.<br/>
          <br/>
            <i>Format</i>: int64<br/>
            <i>Minimum</i>: 0<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### TestRun.status.environmentRequests[index]
<sup><sup>[↩ Parent](#testrunstatus)</sup></sup>



ObjectReference contains enough information to let you inspect or modify the referred object.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>apiVersion</b></td>
        <td>string</td>
        <td>
          API version of the referent.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>fieldPath</b></td>
        <td>string</td>
        <td>
          If referring to a piece of an object instead of an entire object, this string
should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
For example, if the object reference is to a container within a pod, this would take on a value like:
"spec.containers{name}" (where "name" refers to the name of the container that triggered
the event) or if no container name is specified "spec.containers[2]" (container with
index 2 in this pod). This syntax is chosen only to have some well-defined way of
referencing a part of an object.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>kind</b></td>
        <td>string</td>
        <td>
          Kind of the referent.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>namespace</b></td>
        <td>string</td>
        <td>
          Namespace of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>resourceVersion</b></td>
        <td>string</td>
        <td>
          Specific resourceVersion to which this reference is made, if any.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>uid</b></td>
        <td>string</td>
        <td>
          UID of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### TestRun.status.suiteRunners[index]
<sup><sup>[↩ Parent](#testrunstatus)</sup></sup>



ObjectReference contains enough information to let you inspect or modify the referred object.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>apiVersion</b></td>
        <td>string</td>
        <td>
          API version of the referent.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>fieldPath</b></td>
        <td>string</td>
        <td>
          If referring to a piece of an object instead of an entire object, this string
should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
For example, if the object reference is to a container within a pod, this would take on a value like:
"spec.containers{name}" (where "name" refers to the name of the container that triggered
the event) or if no container name is specified "spec.containers[2]" (container with
index 2 in this pod). This syntax is chosen only to have some well-defined way of
referencing a part of an object.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>kind</b></td>
        <td>string</td>
        <td>
          Kind of the referent.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>namespace</b></td>
        <td>string</td>
        <td>
          Namespace of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>resourceVersion</b></td>
        <td>string</td>
        <td>
          Specific resourceVersion to which this reference is made, if any.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>uid</b></td>
        <td>string</td>
        <td>
          UID of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>

# etos.eiffel-community.github.io/v1alpha2

Resource Types:

- [Iut](#iut)




## Iut
<sup><sup>[↩ Parent](#etoseiffel-communitygithubiov1alpha2 )</sup></sup>






Iut is the Schema for the iuts API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>etos.eiffel-community.github.io/v1alpha2</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>Iut</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#iutspec">spec</a></b></td>
        <td>object</td>
        <td>
          spec defines the desired state of Iut<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#iutstatus">status</a></b></td>
        <td>object</td>
        <td>
          status defines the observed state of Iut<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Iut.spec
<sup><sup>[↩ Parent](#iut)</sup></sup>



spec defines the desired state of Iut

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>environmentRequest</b></td>
        <td>string</td>
        <td>
          EnvironmentRequest is the name of the environmentrequest which requested this IUT.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td>
          ID is the ID for the IUT. The ID is a UUID, any version, and regex matches that.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>identity</b></td>
        <td>string</td>
        <td>
          Identity is the PackageURL definition of the IUT.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>provider_id</b></td>
        <td>string</td>
        <td>
          ProviderID is the name of the Provider used to create this Iut.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>provider_data</b></td>
        <td>JSON</td>
        <td>
          ProviderData is specific data provided by the IUT providers<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Iut.status
<sup><sup>[↩ Parent](#iut)</sup></sup>



status defines the observed state of Iut

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>completionTime</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#iutstatusconditionsindex">conditions</a></b></td>
        <td>[]object</td>
        <td>
          conditions represent the current state of the Iut resource.
Each condition has a unique type and reflects the status of a specific aspect of the resource.

Standard condition types include:
- "Available": the resource is fully functional
- "Progressing": the resource is being created or updated
- "Degraded": the resource failed to reach or maintain its desired state

The status of each condition is one of True, False, or Unknown.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Iut.status.conditions[index]
<sup><sup>[↩ Parent](#iutstatus)</sup></sup>



Condition contains details for one aspect of the current state of this API Resource.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>lastTransitionTime</b></td>
        <td>string</td>
        <td>
          lastTransitionTime is the last time the condition transitioned from one status to another.
This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.<br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>message</b></td>
        <td>string</td>
        <td>
          message is a human readable message indicating details about the transition.
This may be an empty string.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>reason</b></td>
        <td>string</td>
        <td>
          reason contains a programmatic identifier indicating the reason for the condition's last transition.
Producers of specific condition types may define expected values and meanings for this field,
and whether the values are considered a guaranteed API.
The value should be a CamelCase string.
This field may not be empty.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>status</b></td>
        <td>enum</td>
        <td>
          status of the condition, one of True, False, Unknown.<br/>
          <br/>
            <i>Enum</i>: True, False, Unknown<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          type of condition in CamelCase or in foo.example.com/CamelCase.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>observedGeneration</b></td>
        <td>integer</td>
        <td>
          observedGeneration represents the .metadata.generation that the condition was set based upon.
For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date
with respect to the current state of the instance.<br/>
          <br/>
            <i>Format</i>: int64<br/>
            <i>Minimum</i>: 0<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>
