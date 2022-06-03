|Event Type|Description|
|---|---|
|`session.start`|SessionEvent indicates that session has been initiated
or updated by a joining party on the server|
|`session.end`|`session.end` indicates that a session has ended|
|`session.upload`|`session.upload` indicates that session has been uploaded to the external storage|
|`session.join`|`session.join` indicates that someone joined a session|
|`session.leave`|`session.leave` indicates that someone left a session|
|`session.data`|Data transfer events.|
|`client.disconnect`|`client.disconnect` is emitted when client is disconnected
by the server due to inactivity or any other reason|
|`user.login`|`user.login` indicates that a user logged into web UI or via tsh|
|`user.update`|`user.update` is emitted when the user is updated.|
|`user.delete`|`user.delete` is emitted when the user is deleted.|
|`user.create`|`user.create` is emitted when the user is created.|
|`user.password_change`|`user.password_change` is when the user changes their own password.|
|`access_request.create`|`access_request.create` is emitted when a new access request is created.|
|`access_request.update`|`access_request.update` is emitted when a request's state is updated.|
|`access_request.review`|`access_request.review` is emitted when a review is applied to a request.|
|`access_request.delete`|`access_request.delete` is emitted when a new access request is deleted.|
|`access_request.search`|`access_request.search` is emitted when a user searches for
resources as part of a search-based access request.|
|`recovery_token.create`|`recovery_token.create` is emitted when a new recovery token is created.|
|`reset_password_token.create`|`reset_password_token.create` is emitted when a new reset password token is created.|
|`privilege_token.create`|`privilege_token.create` is emitted when a new user privilege token is created.|
|`exec`|`exec` is an exec command executed by script or user on
the server side|
|`subsystem`|`subsystem` is the result of the execution of a subsystem.|
|`x11-forward`|X11 forwarding event|
|`port`|Port forwarding event|
|`auth`|`auth` is authentication attempt that either
succeeded or failed based on event status|
|`scp`|`scp` means data transfer that occurred on the server|
|`sftp`|`sftp` means a user attempted a file operation|
|`resize`|`resize` means that some user resized PTY on the client|
|`session.command`|`session.command` is emitted when an executable is run within a session.|
|`session.disk`|`session.disk` is emitted when a file is opened within an session.|
|`session.network`|`session.network` is emitted when a network connection is initiated with a
session.|
|`role.created`|`role.created` fires when role is created/updated.|
|`role.deleted`|`role.deleted` fires when role is deleted.|
|`trusted_cluster.create`|`trusted_cluster.create` is the event for creating a trusted cluster.|
|`trusted_cluster.delete`|`trusted_cluster.delete` is the event for removing a trusted cluster.|
|`trusted_cluster_token.create`|`trusted_cluster_token.create` is the event for
creating new join token for a trusted cluster.|
|`github.created`|`github.created` fires when a Github connector is created/updated.|
|`github.deleted`|`github.deleted` fires when a Github connector is deleted.|
|`oidc.created`|`oidc.created` fires when OIDC connector is created/updated.|
|`oidc.deleted`|`oidc.deleted` fires when OIDC connector is deleted.|
|`saml.created`|`saml.created` fires when SAML connector is created/updated.|
|`saml.deleted`|`saml.deleted` fires when SAML connector is deleted.|
|`session.rejected`|SessionRejected fires when a user's attempt to create an authenticated
session has been rejected due to exceeding a session control limit.|
|`session.connect`|SessionConnect is emitted when any ssh connection is made|
|`app.create`|`app.create` is emitted when an application resource is created.|
|`app.update`|`app.update` is emitted when an application resource is updated.|
|`app.delete`|`app.delete` is emitted when an application resource is deleted.|
|`app.session.start`|`app.session.start` is emitted when a user is issued an application certificate.|
|`app.session.end`|`app.session.end` is emitted when a user connects to a TCP application.|
|`app.session.chunk`|`app.session.chunk` is emitted at the start of a 5 minute chunk on each
proxy. This chunk is used to buffer 5 minutes of audit events at a time
for applications.|
|`app.session.request`|`app.session.request` is an HTTP request and response.|
|`db.create`|`db.create` is emitted when a database resource is created.|
|`db.update`|`db.update` is emitted when a database resource is updated.|
|`db.delete`|`db.delete` is emitted when a database resource is deleted.|
|`db.session.query`|`db.session.query` is emitted when a database client executes
a query.|
|`kube.request`|`kube.request` fires when a proxy handles a generic kubernetes
request.|
|`kube.create`|`kube.create` is emitted when a kubernetes cluster resource is created.|
|`kube.update`|`kube.update` is emitted when a kubernetes cluster resource is updated.|
|`kube.delete`|`kube.delete` is emitted when a kubernetes cluster resource is deleted.|
|`mfa.add`|`mfa.add` is an event type for users adding MFA devices.|
|`mfa.delete`|`mfa.delete` is an event type for users deleting MFA devices.|
|`lock.created`|`lock.created` fires when a lock is created/updated.|
|`lock.deleted`|`lock.deleted` fires when a lock is deleted.|
|`recovery_code.generated`|`recovery_code.generated` is an event type for generating a user's recovery tokens.|
|`recovery_code.used`|`recovery_code.used` is an event type when a recovery token was used.|
|`windows.desktop.session.end`|`windows.desktop.session.end` is emitted when a user disconnects
from a desktop.|
|`cert.create`|`cert.create` is emitted when a certificate is issued.|
|`cert.generation_mismatch`|`cert.generation_mismatch` is emitted when a renewable
certificate's generation counter is invalid.|
|`session.recording.access`|`session.recording.access` is emitted when a session recording is accessed|
