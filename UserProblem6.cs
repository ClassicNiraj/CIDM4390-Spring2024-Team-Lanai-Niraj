public class PerformanceAlertTests
{
    [Fact]
    public async Task Admin_Should_ReceiveAlert_When_ResponseTimeIncreases()
    {
        // Arrange: Initialize the performance monitoring and alert services
        var performanceMonitoringService = new PerformanceMonitoringService(); // Service that monitors application performance
        var alertService = new AlertService();
        var admin = new User { Name = "AdminUser" }; 
        
        // Simulate monitoring response times and detecting an increase
        var increasedResponseTimeDetected = performanceMonitoringService.SimulateResponseTimeIncrease();

        // Act: Send an alert if an increase in response time is detected
        bool alertSent = false;
        if (increasedResponseTimeDetected)
        {
            alertSent = await alertService.SendAlertToUser(admin, "Increase in response time detected.");
        }

        // Assert: Verify that the admin receives an alert when there is an increase in response time
        Assert.True(alertSent, "Admin should receive an alert for increased response time.");
    }
}

// Mock classes for demonstration purposes
public class PerformanceMonitoringService
{
    public bool SimulateResponseTimeIncrease()
    {
        return true; // Simulates detection of increased response time
    }
}

public class AlertService
{
    public async Task<bool> SendAlertToUser(User user, string message)
    {
        return await Task.FromResult(true); // Simulates successful alert sending
    }
}

public class User
{
    public string Name { get; set; }
}