public class QueryJobTests
{
    [Fact]
    public void Admin_Should_BeAbleToSpecifyRules_AndRefineScrapingResults()
    {
        // Arrange: Initialize the scraping service and rule configuration
        var scrapingService = new ScrapingService();
        var ruleService = new RuleService();

        // Simulate defining a rule for scraping
        var rule = new Rule { Criteria = "Specific Criteria", Priority = "High" };
        ruleService.AddRule(rule);

        // Act: Apply the rule and refine scraping results
        var results = scrapingService.RefineScrapingResultsWithRules(rule);

        // Assert: Verify that results are refined based on the specified rule
        Assert.NotEmpty(results);
        Assert.All(results, result => Assert.Contains(rule.Criteria, result.Content));
    }
}

// Mock classes for demonstration purposes
public class ScrapingService
{
    public Result[] RefineScrapingResultsWithRules(Rule rule)
    {
        return new[]
        {
            new Result { Content = "Result matching Specific Criteria" },
            new Result { Content = "Another result matching Specific Criteria" }
        };
    }
}

public class RuleService
{
    public void AddRule(Rule rule)
    {
        // This method simulates adding a new rule for refining scraping results
    }
}

public class Rule
{
    public string Criteria { get; set; }
    public string Priority { get; set; }
}

public class Result
{
    public string Content { get; set; }
}