document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('analyzeForm');
    const urlInput = document.getElementById('url');
    const analyzeBtn = document.querySelector('.analyze-btn');
    const btnText = document.querySelector('.btn-text');
    const btnLoading = document.querySelector('.btn-loading');
    const resultsDiv = document.getElementById('results');
    const resultsContent = document.getElementById('resultsContent');
    const errorDiv = document.getElementById('error');
    const errorContent = document.getElementById('errorContent');

    form.addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const url = urlInput.value.trim();
        if (!url) {
            showError('Enter a URL to analyze');
            return;
        }

        setLoadingState(true);
        hideResults();
        hideError();

        try {
            const response = await fetch('/analyze', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: `url=${encodeURIComponent(url)}`
            });

            const data = await response.json();

            if (!response.ok) {
                if (data.status_code) {
                    showError(`HTTP ${data.status_code}: ${data.description}`);
                } else {
                    showError(data.error || 'Error occurred while analyzing the URL');
                }
                return;
            }

            displayResults(data);
            
        } catch (error) {
            console.error('Error:', error);
            showError('Failed to connect to the server');
        } finally {
            setLoadingState(false);
        }
    });

    function setLoadingState(loading) {
        if (loading) {
            btnText.style.display = 'none';
            btnLoading.style.display = 'inline-flex';
            analyzeBtn.disabled = true;
        } else {
            btnText.style.display = 'inline';
            btnLoading.style.display = 'none';
            analyzeBtn.disabled = false;
        }
    }

    function displayResults(data) {
        const html = `
            <div class="result-item">
                <span class="result-label">URL:</span>
                <span class="result-value">${escapeHtml(data.url)}</span>
            </div>
            <div class="result-item">
                <span class="result-label">HTML Version:</span>
                <span class="result-value">${escapeHtml(data.html_version)}</span>
            </div>
            <div class="result-item">
                <span class="result-label">Page Title:</span>
                <span class="result-value">${escapeHtml(data.title || 'No title found')}</span>
            </div>
            <div class="result-item">
                <span class="result-label">Headings:</span>
                <span class="result-value">
                    ${formatHeadings(data.headings)}
                </span>
            </div>
            <div class="result-item">
                <span class="result-label">Internal Links:</span>
                <span class="result-value">${data.internal_links}</span>
            </div>
            <div class="result-item">
                <span class="result-label">External Links:</span>
                <span class="result-value">${data.external_links}</span>
            </div>
            <div class="result-item">
                <span class="result-label">Inaccessible Links:</span>
                <span class="result-value">${data.inaccessible_links}</span>
            </div>
            <div class="result-item">
                <span class="result-label">Login Form:</span>
                <span class="result-value">${data.has_login_form ? 'Yes' : 'No'}</span>
            </div>
            <div class="result-item">
                <span class="result-label">Analysis Time:</span>
                <span class="result-value">${new Date(data.analysis_time).toLocaleString()}</span>
            </div>
        `;

        resultsContent.innerHTML = html;
        resultsDiv.style.display = 'block';
        
        resultsDiv.scrollIntoView({ behavior: 'smooth' });
    }

    function formatHeadings(headings) {
        if (!headings || Object.keys(headings).length === 0) {
            return 'No headings found';
        }

        const headingItems = Object.entries(headings)
            .map(([level, count]) => `<span class="heading-item">${level}: ${count}</span>`)
            .join('');

        return `<div class="headings-grid">${headingItems}</div>`;
    }

    function showError(message) {
        errorContent.textContent = message;
        errorDiv.style.display = 'block';
        errorDiv.scrollIntoView({ behavior: 'smooth' });
    }

    function hideResults() {
        resultsDiv.style.display = 'none';
    }

    function hideError() {
        errorDiv.style.display = 'none';
    }


    function escapeHtml(text) {
        const div = document.createElement('div');
        div.textContent = text;
        return div.innerHTML;
    }

    urlInput.addEventListener('input', function() {
        const url = this.value.trim();
        if (url && !isValidURL(url)) {
            this.setCustomValidity('Please enter a valid URL');
        } else {
            this.setCustomValidity('');
        }
    });

    function isValidURL(string) {
        try {
            new URL(string);
            return true;
        } catch (_) {
            try {
                new URL('https://' + string);
                return true;
            } catch (_) {
                return false;
            }
        }
    }
});
