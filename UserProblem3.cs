public class NotificationTests
{
    [Fact]
    public async Task Admin_Should_ReceiveNotification_OnCriticalEvent()
    {
        // Arrange: Initialize the monitoring and notification services
        var monitoringService = new MonitoringService(); 
        var notificationService = new NotificationService(); 
        var admin = new User { Name = "AdminUser" }; 

        // Act: Simulate a critical event and attempt to send a notification to the admin
        var criticalEvent = monitoringService.SimulateCriticalEvent();
        var notificationSent = await notificationService.SendNotificationToUser(admin, criticalEvent);

        // Assert: Verify that the admin receives a notification regarding the critical event
        Assert.True(notificationSent, "Admin should receive a notification for the critical event.");
    }
}

// Mock classes for demonstration purposes
public class MonitoringService
{
    public CriticalEvent SimulateCriticalEvent()
    {
        return new CriticalEvent { IsCritical = true, Message = "Critical system failure detected." };
    }
}

public class NotificationService
{
    public async Task<bool> SendNotificationToUser(User user, CriticalEvent criticalEvent)
    {
        return await Task.FromResult(criticalEvent.IsCritical);
    }
}

public class User
{
    public string Name { get; set; }
}

public class CriticalEvent
{
    public bool IsCritical { get; set; }
    public string Message { get; set; }
}